package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_Vault_NewConfig(t *testing.T) {
	err := env.Save(envVars...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			wantErr:   true,
			errWanted: "missing either ROLE or GITHUB_TOKEN env var",
		}, {
			name:     "github environment variables",
			envslice: []string{"GITHUB_TOKEN", mock.Token},
			wantErr:  false,
		}, {
			name:     "k8 environment variables",
			envslice: []string{"MOUNT_PATH", mock.Path, "SERVICE_ACCOUNT_PATH", mock.Path, "ROLE", mock.Role},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnv(t, tt.envslice)

			vault := &Vault{}
			err = vault.Configure()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
