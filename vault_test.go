package vault

import (
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {

	//Test broken config
	vault, err := New()
	if err.Error() != "missing MOUNT_PATH" {
		t.Errorf("Expected error missing MOUNT_PATH, got: %v", err)
	}

	os.Setenv("GITHUB_TOKEN", "no_token")

	//Test broken client
	os.Setenv("VAULT_CACERT", "noFile.tull")
	vault, err = New()
	if !strings.Contains(err.Error(), "while getting CA Certs: failed to read CA file") {
		t.Errorf("Expected file not found, got: %v", err)
	}

	os.Setenv("VAULT_CACERT", "")

	//Test broken authentification
	vault, err = New()
	if !strings.Contains(err.Error(), "while do-ing http request: Post https://127.0.0.1:8200/v1/auth/github/login: dial tcp 127.0.0.1:8200:") {
		t.Errorf("Expected rejected connection, got: %v", err)
	}

	//TODO: need valid test-token
	//Test successful creation
	//vault, err = New()
	//if err != nil {
	//	t.Error(err)
	//}
	if vault.Config.VaultAddr != "https://127.0.0.1:8200" {
		t.Errorf("Error in configuration, unexpected address: %v", err)
	}
	if vault.Client.HTTP == nil {
		t.Error("Error in client, missing HTTP client")
	}
	//if vault.Token.Auth.ClientToken == "" {
	//	t.Error("Error in Token, missing Auth.ClientToken")
	//}
}
