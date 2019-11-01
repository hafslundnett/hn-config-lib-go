package vault

import (
	"os"
	"strings"
	"testing"
)

// makeURL returns a correctly formatted url for Vault http requests based on conffiguration and internal path
func makeURL(cfg Config, path string) string {
	version := "/v1/"
	address := cfg.Addr

	return address + version + path
}

// SetEnv sets environment variables for creating a Vault client
func SetEnv(vaultAddr, pemCert, githubToken, k8ServicePath, k8MountPath, k8Role string) {
	os.Setenv("VAULT_ADDR", vaultAddr)
	os.Setenv("VAULT_CACERT", pemCert)
	os.Setenv("GITHUB_TOKEN", githubToken)
	os.Setenv("SERVICE_ACCOUNT_PATH", k8ServicePath)
	os.Setenv("MOUNT_PATH", k8MountPath)
	os.Setenv("ROLE", k8Role)
}

// Testing helpers

// testing constants, used for consistency across tests
const (
	mockAddr  = "mock.addr"
	mockToken = "mock_token"
	mockRole  = "mock_role"
	mockFile  = "mock.file"
	mockPath  = "mock/path"
)

// assertNoErr if you don't expect an error
func assertNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, got: %v", err)
	}
}

// assertErr if you expect an error
func assertErr(t *testing.T, err error, errSubstr string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error, got none")
	} else if !strings.Contains(err.Error(), errSubstr) {
		t.Fatalf("Unexpected error, got: %v, wanted error containing substing: %s", err, errSubstr)
	}
}
