package vault

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_GetSecret(t *testing.T) {
	vault, err := New()
	assert.NoErr(t, err)

	// Get secret outside policy area
	secret, err := vault.GetSecret(mock.Path)
	assert.Err(t, err, "403 Forbidden")

	if secret != nil {
		t.Error("secret should be nil")
	}

	// Get secret within policy area
	secretTestPath := os.Getenv("SECRET_TEST_PATH")
	secret, err = vault.GetSecret(secretTestPath)
	assert.NoErr(t, err)

	if secret == nil {
		t.Error("secret is empty")
	}
}
