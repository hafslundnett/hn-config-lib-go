package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_Authenticate(t *testing.T) {
	vault := Vault{}

	err := vault.NewConfig()
	assert.NoErr(t, err)

	err = vault.NewClient()
	assert.NoErr(t, err)

	// Test with valid test token
	err = vault.Authenticate()
	assert.NoErr(t, err)

	// Test with invalid github token
	vault.GithubToken = mock.Token
	err = vault.Authenticate()
	assert.Err(t, err, "http error, status 500")

	// Test with invalid k8 token
	vault.GithubToken = ""
	vault.K8Role = mock.Role
	vault.K8MountPath = mock.Path
	vault.K8ServicePath = mock.Path

	err = vault.Authenticate()
	assert.Err(t, err, "while converting token to buffer")
}
