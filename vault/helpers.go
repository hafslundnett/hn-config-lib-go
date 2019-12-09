package vault

var envVars = []string{
	"VAULT_ADDR",
	"VAULT_CACERT",
	"GITHUB_TOKEN",
	"SERVICE_ACCOUNT_PATH",
	"MOUNT_PATH",
	"ROLE",
}

// makeURL returns a correctly formatted url for Vault http requests
func makeURL(address, path string) string {
	version := "/v1/"
	return address + version + path
}
