package hid

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_HID_GetToken(t *testing.T) {
	type args struct {
		user   string
		secret string
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantErr   bool
		errWanted string
	}{
		{
			name:      "mock credentials",
			args:      args{mock.ID, mock.Secret},
			wantErr:   true,
			errWanted: "400 Bad Request",
		}, {
			name:    "test credentials",
			args:    args{os.Getenv("TEST_HID_ID"), os.Getenv("TEST_HID_SECRET")},
			want:    3600,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hid, err := New()
			assert.NoErr(t, err)

			gotToken, err := hid.GetToken(tt.args.user, tt.args.secret)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr {
				got := gotToken.Exp
				assert.Result(t, got, tt.want)
			}
		})
	}
}
