package vault

import (
	"testing"
)

func TestNew(t *testing.T) {
	SetEnv("", "", "")

	//Test broken config
	vault, err := New()
	assertErr(t, err, "missing MOUNT_PATH")

	//Test broken client
	SetEnv("", mockFile, mockToken)
	vault, err = New()
	assertErr(t, err, "while getting CA Certs: failed to read CA file")

	//Test broken authentification
	SetEnv("", "", mockToken)
	vault, err = New()
	assertErr(t, err, "while do-ing http request: Post https://127.0.0.1:8200/v1/auth/github/login: dial tcp 127.0.0.1:8200:")
	
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
