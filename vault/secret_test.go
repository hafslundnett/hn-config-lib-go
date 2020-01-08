package vault

import (
	"net/http"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func TestVault_GetSecret(t *testing.T) {
	tests := []struct {
		name      string
		vault     Vault
		want      *Secret
		wantErr   bool
		errWanted string
	}{
		{
			name:      "forbidden access",
			vault:     Vault{Config: Config{Client: mock.ClientForbidden}},
			wantErr:   true,
			errWanted: "while renewing credentials",
		}, {
			name:    "access granted",
			vault:   Vault{Config: Config{Client: mock.Client}},
			want:    &Secret{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.vault.GetSecret("")
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			assert.DeepResult(t, got, tt.want)
		})
	}
}

func Test_secretsReq(t *testing.T) {
	type args struct {
		url  string
		auth string
	}
	tests := []struct {
		name string
		args args
		want *http.Request
	}{
		{
			name: "build request",
			args: args{mock.Addr, mock.Token},
			want: mock.Request(t, "GET", mock.Addr, "", "X-Vault-Token", mock.Token),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := secretsReq(tt.args.url, tt.args.auth)

			assert.NoErr(t, err)

			assert.Result(t, got.Method, tt.want.Method)
			assert.DeepResult(t, got.URL, tt.want.URL)
			assert.DeepResult(t, got.Header, tt.want.Header)
		})
	}
}

func TestVault_do(t *testing.T) {
	tests := []struct {
		name      string
		vault     Vault
		wantErr   bool
		errWanted string
	}{
		{
			name:      "forbidden access, without authentication",
			vault:     Vault{Config: Config{Client: mock.ClientForbidden}},
			wantErr:   true,
			errWanted: "while renewing credentials",
		}, {
			name:    "access granted",
			vault:   Vault{Config: Config{Client: mock.Client}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.vault.do(nil, nil)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}
}
