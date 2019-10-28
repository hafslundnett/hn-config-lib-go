package vault

import (
	"testing"
)

func Test_Authenticate(t *testing.T) {
	SetEnv("", "", mockToken)
	vault := Vault{}

	err := vault.NewConfig()
	assertNoErr(t, err)

	err = vault.NewClient()
	assertNoErr(t, err)

	//Test with github token
	err = vault.Authenticate()
	assertErr(t, err, "while do-ing http request:")

	//Test with k8 token
	vault.Config.GitHubToken = ""
	vault.Config.K8Token = mockToken
	vault.Config.Role = mockRole

	err = vault.Authenticate()
	assertErr(t, err, "while do-ing http request:")
}
