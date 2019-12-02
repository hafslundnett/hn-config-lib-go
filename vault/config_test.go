package vault

import (
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
	"testing"
)

func Test_NewConfig(t *testing.T) { //TODO: test kubernetes configuration
	setEnv()
	vault := Vault{}

	//Test with no environment variables
	err := vault.NewConfig()
	assert.Err(t, err, "missing ROLE env var")

	//Test with environment variables
	setEnv("GITHUB_TOKEN", mock.Token)
	err = vault.NewConfig()
	assert.NoErr(t, err)

	//Test with k8 environment variables
	setEnv("MOUNT_PATH", mock.Path, "SERVICE_ACCOUNT_PATH", mock.Path, "ROLE", mock.Role)
	err = vault.NewConfig()
	assert.NoErr(t, err)
}
