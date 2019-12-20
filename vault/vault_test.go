package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

var envVars = []string{
	"VAULT_ADDR",
	"VAULT_CACERT",
	"GITHUB_TOKEN",
	"SERVICE_ACCOUNT_PATH",
	"MOUNT_PATH",
	"ROLE",
}

// replaceEnv compacts environment variables handling to increase readability of tests.
func replaceEnv(t *testing.T, vars []string) {
	t.Helper()
	err := env.Clear(envVars...)
	assert.NoErr(t, err)
	err = env.Set(vars...)
	assert.NoErr(t, err)
}

func Test_New(t *testing.T) {
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
			envslice:  []string{},
			wantErr:   true,
			errWanted: "missing either ROLE or GITHUB_TOKEN env var",
		}, {
			name:      "broken authentification",
			envslice:  []string{"GITHUB_TOKEN", mock.Token},
			wantErr:   true,
			errWanted: "while do-ing http request",
		}, {
			name:      "broken client",
			envslice:  []string{"GITHUB_TOKEN", mock.Token, "VAULT_CACERT", mock.File},
			wantErr:   true,
			errWanted: "while loading CA cert from file",
		}, {
			name: "successful creation",
			// default environment
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envslice != nil {
				replaceEnv(t, tt.envslice)
			}

			vault, err := New()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr && vault.Token.Auth.Accessor == "" {
				t.Error("No token recieved")
			}

			err = env.Reset()
			assert.NoErr(t, err)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}

func Test_Vault_NewClient(t *testing.T) {
	tests := []struct {
		name      string
		vault     *Vault
		wantErr   bool
		errWanted string
	}{
		{
			name:      "pool error",
			vault:     &Vault{Config: Config{PemCert: mock.File}},
			wantErr:   true,
			errWanted: "failed to read CA file",
		}, {
			name:    "functional pool",
			vault:   &Vault{Config: Config{PemCert: ""}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.vault.NewClient()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr && tt.vault.Client == nil {
				t.Error("No http client recieved")
			}
		})
	}
}
