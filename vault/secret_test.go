package vault

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func Test_GetSecret(t *testing.T) {
	setEnv("GITHUB_TOKEN", mock.Token)

	vault, err := New()
	assert.Err(t, err, "while do-ing http request:")

	_, err = vault.GetSecret(mock.Path)
	assert.Err(t, err, "while do-ing http request:")
}
