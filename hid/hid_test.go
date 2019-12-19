package hid

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

func Test_New(t *testing.T) {
	// test with defaults. Fails successfully if HID is not running locally
	env.Set("HID_ADDR", "") //TODO: err
	testHID, err := New()
	assert.Err(t, err, "while renewing HID public key set")

	// test against HID test server
	addr := "https://hid.test-hafslundnett.io"
	env.Set("HID_ADDR", addr)

	testHID, err = New()
	assert.NoErr(t, err)

	assert.Result(t, testHID.Addr, addr)
	assert.Result(t, testHID.JWKSuri, addr+"/.well-known/openid-configuration/jwks")
	//TODO: reset env
}
