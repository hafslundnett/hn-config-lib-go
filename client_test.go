package vault

import (
	"net/http"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	vault := Vault{}

	os.Setenv("VAULT_CACERT", "mock.file")
	os.Setenv("GITHUB_TOKEN", "mock_token")

	err := vault.NewConfig()
	if err != nil {
		t.Errorf("Unexpected error, got: %v", err)
	}

	//Test bad pool
	err = vault.NewClient()
	if err == nil {
		t.Error("Expected error, got none")
	}

	vault.Config.PemCert = ""

	//Test all good
	err = vault.NewClient()
	if err != nil {
		t.Errorf("Unexpected error, got: %v", err)
	}
}

func TestRemoteCall(t *testing.T) {
	vault := Vault{}

	os.Setenv("GITHUB_TOKEN", "mock_token")

	err := vault.NewConfig()
	if err != nil {
		t.Errorf("Unexpected error, got: %v", err)
	}

	err = vault.NewClient()
	if err != nil {
		t.Errorf("Unexpected error, got: %v", err)
	}

	req := new(http.Request)

	//Test bad call
	err = vault.Client.remoteCall(req, &vault.Token)
	if err == nil {
		t.Error("Expected error, got none")
	}

	//Test wrong resp code
	//TODO: need test env

	//Test all good
	//TODO: need test env
}

func TestMakeUrl(t *testing.T) {
	//vault := Vault{}

	//Test no config
	//makeURL()

	//Test all good
}
