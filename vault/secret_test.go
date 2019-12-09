package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_GetSecret(t *testing.T) {
	vault, err := New()
	assert.NoErr(t, err)

	secret, err := vault.GetSecret(mock.Path)
	assert.Err(t, err, "403 Forbidden")

	if secret != nil { // Better test when available
		t.Error("secret should be nil")
	}
}
