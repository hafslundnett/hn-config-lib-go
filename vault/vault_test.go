package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_New(t *testing.T) {
	setEnv()

	// Test broken config
	vault, err := New()
	assert.Err(t, err, "missing ROLE env var")

	// Test broken client
	setEnv("VAULT_CACERT", mock.File, "GITHUB_TOKEN", mock.Token)
	vault, err = New()
	assert.Err(t, err, "while loading CA cert from file: failed to read CA file")

	// Test broken authentification
	setEnv("GITHUB_TOKEN", mock.Token)
	vault, err = New()
	assert.Err(t, err, "while do-ing http request: Post https://127.0.0.1:8200/v1/auth/github/login: dial tcp 127.0.0.1:8200:")

	// TODO: need valid test-token
	// Test successful creation
	//vault, err = New()
	//if err != nil {
	//	t.Error(err)
	//}
	if vault.Config.Addr != "https://127.0.0.1:8200" {
		t.Errorf("Error in configuration, unexpected address: %v", err)
	}
	if vault.Client.HTTP == nil {
		t.Error("Error in client, missing HTTP client")
	}
	//if vault.Token.Auth.ClientToken == "" {
	//	t.Error("Error in Token, missing Auth.ClientToken")
	//}
}

func Test_NewClient(t *testing.T) {
	setEnv("VAULT_CACERT", mock.File, "GITHUB_TOKEN", mock.Token)
	vault := Vault{}

	err := vault.NewConfig()
	assert.NoErr(t, err)

	// Test with pool error
	err = vault.NewClient()
	assert.Err(t, err, "while loading CA cert from file: failed to read CA file")

	// Test with functional pool
	vault.PemCert = ""
	err = vault.NewClient()
	assert.NoErr(t, err)
}
