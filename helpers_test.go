package vault

import (
	"os"
	"testing"
)

func Test_makeUrl(t *testing.T) {
	vault := Vault{}

	// Test empty config
	path := ""
	want := "/v1/"
	got := makeURL(vault.Config, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}

	// Test good config
	vault.Config.VaultAddr = mockAddr
	path = mockPath
	want = vault.Config.VaultAddr + "/v1/" + path
	got = makeURL(vault.Config, path)
	if got != want {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}

func Test_SetEnv(t *testing.T) {
	tests := []struct {
		vaultAddr   string
		pemCert     string
		githubToken string
	}{
		{
			vaultAddr:   mockAddr,
			pemCert:     mockFile,
			githubToken: mockToken,
		}, {
			vaultAddr:   "",
			pemCert:     "",
			githubToken: "",
		},
	}

	for _, test := range tests {
		SetEnv(test.vaultAddr, test.pemCert, test.githubToken, "", "", "")

		if os.Getenv("VAULT_ADDR") != test.vaultAddr || os.Getenv("VAULT_CACERT") != test.pemCert || os.Getenv("GITHUB_TOKEN") != test.githubToken {
			t.Fatalf("Unexpected environment variable")
		}
	}
}
