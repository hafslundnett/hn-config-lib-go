package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_New(t *testing.T) {
	env.Save(envVars...)
	env.Clear(envVars...)
	defer env.Reset()

	// Test broken config
	vault, err := New()
	assert.Err(t, err, "missing ROLE env var")

	// Test broken authentification
	env.Set("GITHUB_TOKEN", mock.Token)
	vault, err = New()
	assert.Err(t, err, "while do-ing http request")

	// Test broken client
	env.Set("VAULT_CACERT", mock.File)
	vault, err = New()
	assert.Err(t, err, "while loading CA cert from file")

	// Test successful creation
	env.Reset()
	vault, err = New()
	assert.NoErr(t, err)
	if vault.Token.Auth.Accessor == "" {
		t.Error("No token recieved")
	}
}

func Test_NewClient(t *testing.T) {
	env.Clear(envVars...)
	env.Set("VAULT_CACERT", mock.File, "GITHUB_TOKEN", mock.Token)
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
