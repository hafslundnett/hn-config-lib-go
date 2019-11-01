package vault

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func Test_NewClient(t *testing.T) {
	setEnv("VAULT_CACERT", mock.File, "GITHUB_TOKEN", mock.Token)
	vault := Vault{}

	err := vault.NewConfig()
	assert.NoErr(t, err)

	// Test with pool error
	err = vault.NewClient()
	assert.Err(t, err, "while getting CA Certs: failed to read CA file")

	// Test with functional pool
	vault.PemCert = ""
	err = vault.NewClient()
	assert.NoErr(t, err)
}
