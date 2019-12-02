package hnhttp

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/hafslundnett/hn-config-lib-go/cert"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

// Client expl
type Client struct {
	HTTP *http.Client
}

// NewClient expl
func NewClient(certificates ...string) (*Client, error) {
	pool, err := cert.MakePool(certificates...)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool.Certs,
	}

	transport := &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	client := &Client{
		HTTP: httpClient,
	}

	return client, nil
}

// Do is a wrapper around http.do that error checks and decodes the response.
// dst is an optional destination pointer that must be either implementing io.writer or a struct to be populated by JSON data
// as dst is a variadic argument, multiple dst may be passed, however any additional destinations are ignored with no error.
func (client Client) Do(req *http.Request, dst ...interface{}) error {
	resp, err := client.HTTP.Do(req)
	if err != nil {
		return errors.Wrap(err, "while do-ing http request")
	}

	err = procResp(resp, dst...)
	if err != nil {
		return err
	}

	return nil
}

// Get expl
func (client Client) Get(url string, dst ...interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	err = client.Do(req, dst...)
	if err != nil {
		return err
	}

	return nil
}

// PostForm is a wrapper around http.PostForm that takes multiple kinds of input, and makes error checks and decodes the response.
// src is the source struct or map
// dst is an optional destination pointer that must be either implementing io.writer or a struct to be populated by JSON data
// as dst is a variadic argument, multiple dst may be passed, however any additional destinations are ignored with no error.
func (client Client) PostForm(url string, data map[string][]string, dst ...interface{}) error {
	resp, err := client.HTTP.PostForm(url, data)
	if err != nil {
		return errors.Wrap(err, "while post-ing http request")
	}

	err = procResp(resp, dst...)
	if err != nil {
		return err
	}

	return nil
}

func procResp(resp *http.Response, dst ...interface{}) error {
	//check response code
	if err := checkRespCode(resp); err != nil {
		return err
	}

	// Check if decoding the response is needed.
	if dst != nil {
		if dst[0] != nil {
			defer resp.Body.Close()
			decodeBody(resp.Body, dst[0])
		}
	}

	return nil
}

func decodeBody(body io.ReadCloser, dst interface{}) error {
	// If dst implements io.Writer, write the raw response body.
	if w, ok := dst.(io.Writer); ok {
		_, err := io.Copy(w, body)
		return err
	}

	// Else decode as JSON
	if err := json.NewDecoder(body).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	return nil
}

// checkRespCode checks the status code of an HTTP request.
func checkRespCode(r *http.Response) error {
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return errors.Errorf("http error, status code %d", r.StatusCode)
	}

	return nil
}

/**
func structToValues(i interface{}) (data url.Values) {
	data = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		data.Set(typ.Field(i).Name, fmt.Sprint(iVal.Field(i)))
	}
	return
}
*/
