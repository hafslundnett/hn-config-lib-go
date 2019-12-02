package vault

import (
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
	"testing"
)

func Test_GetSecret(t *testing.T) {
	setEnv("GITHUB_TOKEN", mock.Token)

	vault, err := New()
	assert.Err(t, err, "while do-ing http request:")

	_, err = vault.GetSecret(mock.Path)
	assert.Err(t, err, "while do-ing http request:")
}
