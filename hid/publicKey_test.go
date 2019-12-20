package hid

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

func Test_HID_NewPKS(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		errWanted string
	}{
		{
			name:    "getting public key set",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hid := &HID{}
			err := hid.NewClient()
			assert.NoErr(t, err)
			err = hid.NewConfig()
			assert.NoErr(t, err)

			err = hid.NewPKS()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr && hid.PKS.Keys == nil {
				t.Error("Did not get any keys")
			}
		})
	}
}
