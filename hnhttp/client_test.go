package hnhttp

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"net/http"
	"net/url"
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

	err = client.Do(req)
	assert.Err(t, err, "while do-ing http request:")
}

func Test_PostForm(t *testing.T) {
	client, err := NewClient()
	assert.NoErr(t, err)

	form := url.Values{}
	form.Add("addr", mock.Addr)
	form.Add("path", mock.Path)

	err = client.PostForm(mock.URL, form)
	assert.Err(t, err, "while post-ing http request")
}
