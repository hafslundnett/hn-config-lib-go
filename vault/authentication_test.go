package vault

import (
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
	"testing"
)

func Test_Authenticate(t *testing.T) {
	setEnv("GITHUB_TOKEN", mock.Token)
	vault := Vault{}

	err := vault.NewConfig()
	assert.NoErr(t, err)

	err = vault.NewClient()
	assert.NoErr(t, err)

	// Test with github token
	err = vault.Authenticate()
	assert.Err(t, err, "while do-ing http request:")

	// Test with k8 token
	vault.GithubToken = ""
	vault.K8Role = mock.Role

	err = vault.Authenticate()
	assert.Err(t, err, "while converting token to buffer:")
}
