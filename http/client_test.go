package http

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func Test_NewClient(t *testing.T) {
	// Test with pool error
	_, err := NewClient(mock.File)
	assert.Err(t, err, "while getting CA Certs: failed to read CA file")

	// Test with functional pool
	_, err = NewClient("")
	assert.NoErr(t, err)
}
