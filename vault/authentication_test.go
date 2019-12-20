package vault

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_Vault_Authenticate(t *testing.T) {
	err := env.Save(envVars...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "invalid github token",
			envslice:  []string{"VAULT_ADDR", os.Getenv("VAULT_ADDR"), "GITHUB_TOKEN", mock.Token},
			wantErr:   true,
			errWanted: "http error, status 500",
		}, {
			name:      "invalid k8 token",
			envslice:  []string{"VAULT_ADDR", os.Getenv("VAULT_ADDR"), "MOUNT_PATH", mock.Path, "SERVICE_ACCOUNT_PATH", mock.Path, "ROLE", mock.Role},
			wantErr:   true,
			errWanted: "while converting token to buffer",
		}, {
			name: "valid test token",
			// default environment
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envslice != nil {
				replaceEnv(t, tt.envslice)
			}

			vault := &Vault{}
			err := vault.NewConfig()
			assert.NoErr(t, err)
			err = vault.NewClient()
			assert.NoErr(t, err)

			err = vault.Authenticate()
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

func Test_MakeUrl(t *testing.T) {
	// Test empty
	addr := ""
	path := ""
	want := "/v1/"
	got := makeURL(addr, path)
	assert.Result(t, got, want)

	// Test mock data
	addr = mock.Addr
	path = mock.Path
	want = addr + "/v1/" + path
	got = makeURL(addr, path)
	assert.Result(t, got, want)
}
