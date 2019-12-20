package hid

import (
	"os"
	"testing"
	"time"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_VerifyClaims(t *testing.T) {
	type args struct {
		issuer   string
		audience string
		scope    string
	}
	tests := []struct {
		name      string
		time      time.Time
		args      args
		wantErr   bool
		errWanted string
	}{
		{
			name:      "invalid audience",
			time:      time.Now().Add(time.Minute * 30),
			args:      args{mock.Addr, mock.ID, mock.ID},
			wantErr:   true,
			errWanted: "Invalid audience",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at(tt.time, func() {
				hid, err := New()
				assert.NoErr(t, err)

				token, err := hid.GetToken(os.Getenv("TEST_HID_ID"), os.Getenv("TEST_HID_SECRET"))
				assert.NoErr(t, err)

				jwt, err := hid.Authenticate(token.Raw)
				assert.NoErr(t, err)

				err = VerifyClaims(jwt, tt.args.issuer, tt.args.audience, tt.args.scope)
				assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			})
		})
	}
}
