package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func TestVault_NewConfig(t *testing.T) {
	tests := []struct {
		name      string
		args      map[string]string
		vault     *Vault
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			args:      map[string]string{},
			vault:     &Vault{},
			wantErr:   true,
			errWanted: "missing ROLE env var",
		}, {
			name:    "github environment variables",
			args:    map[string]string{"GITHUB_TOKEN": mock.Token},
			vault:   &Vault{},
			wantErr: false,
		}, {
			name:    "k8 environment variables",
			args:    map[string]string{"MOUNT_PATH": mock.Path, "SERVICE_ACCOUNT_PATH": mock.Path, "ROLE": mock.Role},
			vault:   &Vault{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env.Clear(envVars...)
			env.SetMap(tt.args)

			err := tt.vault.NewConfig()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}
}
