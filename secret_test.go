package vault

import (
	"testing"
)

//temporary example implementation
func TestGetSecret(t *testing.T) {
	vault, err := New()
	if err != nil {
		t.Errorf("while configuring vault: %v", err)
	}

	path := "documentor/kv/storage%2Fdocumentationdata"

	secret, err := vault.GetSecret(path)
	if err != nil {
		t.Errorf("while getting secret: %v", err)
	}

	t.Log(secret)
}
