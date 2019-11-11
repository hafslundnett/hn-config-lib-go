package hid

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

func Test_GetToken(t *testing.T) {
	l := &Login{
		URL:    mock.URL,
		ID:     mock.ID,
		Secret: mock.Path,
	}

	_, err := l.GetToken()
	assert.Err(t, err, "while post-ing http request")
}
