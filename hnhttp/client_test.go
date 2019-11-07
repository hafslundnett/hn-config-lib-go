package hnhttp

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"net/http"
	"testing"
)

//TODO: better tests, need testing server

func Test_NewClient(t *testing.T) {
	// Test with pool error
	_, err := NewClient(mock.File)
	assert.Err(t, err, "while getting CA Certs: failed to read CA file")

	// Test with no pool
	_, err = NewClient()
	assert.NoErr(t, err)
}

func Test_checkRespCode(t *testing.T) {
	resp := &http.Response{
		StatusCode: 404,
	}
	err := checkRespCode(resp)
	assert.Err(t, err, "http error, status code 404")

	resp.StatusCode = 200
	err = checkRespCode(resp)
	assert.NoErr(t, err)
}

func Test_Do(t *testing.T) {
	client, err := NewClient()
	assert.NoErr(t, err)

	req, err := http.NewRequest(http.MethodGet, mock.URL, nil)
	assert.NoErr(t, err)

	_, err = client.Do(req)
	assert.Err(t, err, "while do-ing http request:")
}

func Test_Get(t *testing.T) {
	client, err := NewClient()
	assert.NoErr(t, err)

	req, err := http.NewRequest(http.MethodGet, mock.URL, nil)
	assert.NoErr(t, err)

	err = client.Get(req, nil)
	assert.Err(t, err, "while do-ing http request:")
}
