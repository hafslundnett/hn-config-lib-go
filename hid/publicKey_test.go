package hid

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_HID_newPKS(t *testing.T) {
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
			err := hid.Configure(mock.Client)
			assert.NoErr(t, err)

			err = hid.newPKS()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			//if !tt.wantErr && hid.PKS.Keys == nil {
			//	t.Error("Did not get any keys")
			//}
		})
	}
}
