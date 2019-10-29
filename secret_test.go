package vault

import (
	"testing"
)

func Test_GetSecret(t *testing.T) {
	SetEnv("", "", mockToken, "", "", "")

	vault, err := New()
	assertErr(t, err, "while do-ing http request: ")

	_, err = vault.GetSecret(mockPath)
	assertErr(t, err, "while do-ing http request: Get https://127.0.0.1:8200/v1/mock/path:")
}
