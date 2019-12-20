package hid

import (
	"log"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_HID_NewConfig(t *testing.T) {
	err := env.Save(envVars...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:     "empty environment variables",
			envslice: []string{},
			wantErr:  false,
		}, {
			name:      "bad environment variables",
			envslice:  []string{"HID_ADDR", mock.Addr, "HID_CACERT", mock.File, "HID_DISCOVERY", mock.Path},
			wantErr:   true,
			errWanted: "failed to read CA file \"" + mock.File + "\" from disk",
		}, {
			name:    "good environment variables",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envslice != nil {
				replaceEnv(t, tt.envslice)
			}

			hid := new(HID)
			hid.NewClient()
			err = hid.NewConfig()
			log.Println(hid)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			err = env.Reset()
			assert.NoErr(t, err)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
