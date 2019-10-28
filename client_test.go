package vault

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	SetEnv("", mockFile, mockToken)
	vault := Vault{}

	err := vault.NewConfig()
	assertNoErr(t, err)

	//Test with pool error
	err = vault.NewClient()
	assertErr(t, err, "while getting CA Certs: failed to read CA file")

	//Test with functional pool
	vault.Config.PemCert = ""
	err = vault.NewClient()
	assertNoErr(t, err)
}

func TestRemoteCall(t *testing.T) {
	SetEnv("", "", mockToken)
	vault := Vault{}

	err := vault.NewConfig()
	assertNoErr(t, err)

	err = vault.NewClient()
	assertNoErr(t, err)

	req := new(http.Request)

	//Test bad call
	err = vault.Client.do(req, &vault.Token)
	assertErr(t, err, "while do-ing http request:")

	//Test wrong resp code
	//TODO: need test env

	//Test all good
	//TODO: need test env
}
