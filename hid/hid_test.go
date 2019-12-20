package hid

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

var envVars = []string{
	"HID_ADDR",
	"HID_CACERT",
	"HID_DISCOVERY",
}

// replaceEnv compacts environment variables handling to increase readability of tests.
func replaceEnv(t *testing.T, vars []string) {
	t.Helper()
	err := env.Clear(envVars...)
	assert.NoErr(t, err)
	err = env.Set(vars...)
	assert.NoErr(t, err)
}

func Test_New(t *testing.T) {
	err := env.Save(envVars...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		want      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			envslice:  []string{},
			wantErr:   true, // Fails successfully if HID is not running locally
			errWanted: "while renewing HID public key set",
		}, {
			name: "HID test server",
			want: os.Getenv("HID_ADDR") + "/.well-known/openid-configuration/jwks",
			// default environment
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envslice != nil {
				replaceEnv(t, tt.envslice)
			}

			hid, err := New()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr {
				assert.Result(t, hid.JWKSuri, tt.want)
			}

			err = env.Reset()
			assert.NoErr(t, err)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}

func Test_HID_NewClient(t *testing.T) {
	tests := []struct {
		name      string
		hid       *HID
		wantErr   bool
		errWanted string
	}{
		{
			name:      "pool error",
			hid:       &HID{Config: Config{PemCert: mock.File}},
			wantErr:   true,
			errWanted: "failed to read CA file",
		}, {
			name:    "functional pool",
			hid:     &HID{Config: Config{PemCert: ""}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.hid.NewClient()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr && tt.hid.Client == nil {
				t.Error("No http client recieved")
			}
		})
	}
}
