package testing

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/hid"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

// HIDtoken expl
func HIDtoken(t *testing.T) string {
	hidEndpoint, err := hid.New()
	assert.NoErr(t, err)

	token, err := hidEndpoint.GetToken(os.Getenv("TEST_HID_ID"), os.Getenv("TEST_HID_SECRET"))
	assert.NoErr(t, err)

	return token.Raw
}
