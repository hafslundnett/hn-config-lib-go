package vault

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_Vault_GetSecret(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "get secret outside policy area",
			path:      mock.Path,
			wantErr:   true,
			errWanted: "403 Forbidden",
		}, {
			name:    "get secret within policy area",
			path:    os.Getenv("SECRET_TEST_PATH"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vault, err := New()
			assert.NoErr(t, err)

			secret, err := vault.GetSecret(tt.path)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if tt.wantErr && secret != nil {
				t.Error("secret should be nil")
			}
			if !tt.wantErr && secret == nil {
				t.Error("secret is empty")
			}
		})
	}
}
