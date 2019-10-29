package vault

import (
	"testing"
)

func Test_Authenticate(t *testing.T) {
	SetEnv("", "", mockToken, "", "", "")
	vault := Vault{}

	err := vault.NewConfig()
	assertNoErr(t, err)

	err = vault.NewClient()
	assertNoErr(t, err)

	// Test with github token
	err = vault.Authenticate()
	assertErr(t, err, "while do-ing http request:")

	// Test with k8 token
	vault.Config.GithubToken = ""
	vault.Config.K8Role = mockRole

	err = vault.Authenticate()
	assertErr(t, err, "while converting token to buffer:")
}
