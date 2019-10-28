package vault

import (
	"testing"
)

func TestNewConfig(t *testing.T) { //TODO: test kubernetes configuration
	SetEnv("", "", "")
	vault := Vault{}

	//Test with no environment variables
	err := vault.NewConfig()
	assertErr(t, err, "missing MOUNT_PATH")

	//Test with environment variables
	SetEnv(mockAddr, "", mockToken)

	err = vault.NewConfig()
	assertNoErr(t, err)
}
