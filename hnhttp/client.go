package hnhttp

import (
	"crypto/tls"
	"encoding/json"
	"hafslundnett/x/hn-config-lib/cert"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

// Client contains the http client used to interact with Vault
type Client struct {
	HTTP *http.Client
}

// NewClient returns a http client configured according to the supplied Config, for use with Vault
func NewClient(certificate ...string) (*Client, error) {
	pool, err := cert.MakePool(certificate...)
	if err != nil {
		return nil, errors.Wrap(err, "while getting CA Certs")
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
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

// Get exp
func (client Client) Get(req *http.Request, dst interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Return here if decoding the response isn't needed.
	if dst == nil {
		return nil
	}

	// If dst implements io.Writer, write the raw response body.
	if w, ok := dst.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	// Else decode as JSON
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	return nil
}

// Do is a wrapper around http.do for the Vault http client that error checks and decodes to the destination interface
func (client Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.HTTP.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "while do-ing http request")
	}

	if err := checkRespCode(resp); err != nil {
		return nil, errors.Wrap(err, "while do-ing http request")
	}

	return resp, nil
}

// checkRespCode can be used to check the status code of an HTTP request.
func checkRespCode(r *http.Response) error {
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return errors.Errorf("http error, status code %d", r.StatusCode)
	}

	return nil
}
