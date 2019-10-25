package vault

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) { //TODO: test kubernetes configuration
	vault := Vault{}

	err := vault.NewConfig()
	if err.Error() != "missing MOUNT_PATH" {
		t.Errorf("Expexted error \"missing MOUNT_PATH\", got: %v", err)
	}

	os.Setenv("VAULT_ADDR", "mock.addr")
	os.Setenv("GITHUB_TOKEN", "mock_token")

	err = vault.NewConfig()
	if err != nil {
		t.Error(err)
	}
}
