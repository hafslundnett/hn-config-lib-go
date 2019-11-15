package hid

import (
	"strconv"
	"testing"

	"github.com/giert/silver-octo-packet/assert"
)

func mockIntspec() *intspec {
	return &intspec{
		Active: false,
		Token:  mockToken(),
	}
}

func Test_introspect(t *testing.T) {
	is := mockIntspec()

	valid := is.introspect()
	assert.Result(t, strconv.FormatBool(valid), "false")
}

func Test_remoteIntspec(t *testing.T) {
	hid := mockHID()
	token := mockToken()

	_, err := remoteIntspec(hid, token)
	assert.Err(t, err, "while do-ing http request")
}

func Test_Introspection(t *testing.T) {
	hid := mockHID()
	token := mockToken()

	_, err := hid.Introspection(token)
	assert.Err(t, err, "while do-ing http request")
}
