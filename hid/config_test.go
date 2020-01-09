package hid

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/libhttp"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

var testenv = []string{
	envars["addr"],
	envars["cert"],
	envars["discovery"],
}

func Test_HID_Configure(t *testing.T) {
	err := env.Save(testenv...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		client    libhttp.Client
		want      Config
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			client:    mock.Client,
			want:      Config{},
			wantErr:   true,
			errWanted: "missing env var " + envars["addr"],
		}, {
			name:     "defaulting values",
			envslice: []string{envars["addr"], mock.Addr},
			client:   mock.ClientForbidden,
			want:     Config{mock.Addr, mock.Addr + defJWKS, mock.Addr + defTokenEP, mock.ClientForbidden},
			wantErr:  false,
		}, {
			name:     "with all environment variables",
			envslice: []string{envars["addr"], mock.Addr, envars["discovery"], mock.Path},
			client:   mock.Client,
			want:     Config{Addr: mock.Addr, client: mock.Client},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnv(t, tt.envslice)

			hid := new(HID)
			err = hid.Configure(tt.client)

			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			assert.Result(t, hid.Config, tt.want)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
