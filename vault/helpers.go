package vault

import (
	"fmt"
	"github.com/hafslundnett/hn-config-lib-go/env"
	"math/rand"
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

func getRandomFileName(prefix, suffix string) string {
	i := rand.Intn(100000)
	return fmt.Sprintf("%s_%d.%s", prefix, i, suffix)
}