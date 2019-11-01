package vault

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func Test_MakeUrl(t *testing.T) {
	// Test empty
	addr := ""
	path := ""
	want := "/v1/"
	got := makeURL(addr, path)
	assert.Result(t, got, want)

	// Test mock data
	addr = mock.Addr
	path = mock.Path
	want = addr + "/v1/" + path
	got = makeURL(addr, path)
	assert.Result(t, got, want)
}

/**
func Test_setEnv(t *testing.T) {
	tests := []struct {
		vaultAddr   string
		pemCert     string
		githubToken string
	}{
		{
			vaultAddr:   mock.Addr,
			pemCert:     mock.File,
			githubToken: mock.Token,
		}, {
			vaultAddr:   "",
			pemCert:     "",
			githubToken: "",
		},
	}

	for _, test := range tests {
		setEnv(test.vaultAddr, test.pemCert, test.githubToken)

		if os.Getenv("VAULT_ADDR") != test.vaultAddr || os.Getenv("VAULT_CACERT") != test.pemCert || os.Getenv("GITHUB_TOKEN") != test.githubToken {
			t.Fatalf("Unexpected environment variable")
		}
	}
}
*/
