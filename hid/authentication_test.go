package hid

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func mockToken() Token {
	return Token{
		Access:     mock.Token,
		Expiration: 1000,
	}
}

func Test_GetToken(t *testing.T) {
	hid := &HIDclient{
		Host:   mock.URL,
		Path:   mock.Path,
		Secret: mock.Path,
	}

	_, err := hid.GetToken(mock.ID)
	assert.Err(t, err, "while post-ing http request")
}
