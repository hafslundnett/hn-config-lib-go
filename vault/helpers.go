package vault

import (
	"hafslundnett/x/hn-config-lib/env"
)

// makeURL returns a correctly formatted url for Vault http requests
func makeURL(address, path string) string {
	version := "/v1/"
	return address + version + path
}

// setEnv clears vault env variables, and sets the ones provided. Intended for use in tests
// TODO: better comment
func setEnv(vars ...string) {
	env.SetMap(map[string]string{
		"VAULT_ADDR":           "",
		"VAULT_CACERT":         "",
		"GITHUB_TOKEN":         "",
		"SERVICE_ACCOUNT_PATH": "",
		"MOUNT_PATH":           "",
		"ROLE":                 "",
	})
	env.Set(vars...)
}
