package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_NewConfig(t *testing.T) {
	vault := Vault{}

	//Test with no environment variables
	env.Clear(envVars...)
	err := vault.NewConfig()
	assert.Err(t, err, "missing ROLE env var")

	//Test with environment variables
	env.Set("GITHUB_TOKEN", mock.Token)
	err = vault.NewConfig()
	assert.NoErr(t, err)

	//Test with k8 environment variables
	env.Clear("GITHUB_TOKEN")
	env.Set("MOUNT_PATH", mock.Path, "SERVICE_ACCOUNT_PATH", mock.Path, "ROLE", mock.Role)
	err = vault.NewConfig()
	assert.NoErr(t, err)
}
