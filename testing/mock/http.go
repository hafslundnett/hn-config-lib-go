package mock

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

// Client expl
var Client = client{nil}

// ClientForbidden expl
var ClientForbidden = client{errors.New("http error, status 403: forbidden")}

type client struct {
	err error
}

func (c client) Do(req *http.Request, dst ...interface{}) error {
	return c.err
}

func (c client) Get(url string, dst ...interface{}) error {
	return c.err
}

func (c client) PostForm(url string, data map[string][]string, dst ...interface{}) error {
	return c.err
}

// Body expl
func Body(content string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(content))
}

// Request expl
func Request(t *testing.T, method, url, body string, header ...string) *http.Request {
	t.Helper()

	req, err := http.NewRequest(method, url, Body(body))
	assert.NoErr(t, err)

	for i := 0; i+1 < len(header); i += 2 {
		req.Header.Set(header[i], header[i+1])
	}

	return req
}
