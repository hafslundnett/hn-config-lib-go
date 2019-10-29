package vault

import (
	"testing"
)

func Test_NewConfig(t *testing.T) { //TODO: test kubernetes configuration
	SetEnv("", "", "", "", "", "")
	vault := Vault{}

	//Test with no environment variables
	err := vault.NewConfig()
	assertErr(t, err, "missing ROLE env var")

	//Test with environment variables
	SetEnv("", "", mockToken, "", "", "")
	err = vault.NewConfig()
	assertNoErr(t, err)

	//Test with k8 environment variables
	SetEnv("", "", "", mockPath, mockPath, mockRole)
	err = vault.NewConfig()
	assertNoErr(t, err)
}
