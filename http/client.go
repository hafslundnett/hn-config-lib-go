package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"hafslundnett/x/hn-config-lib/cert"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

// Client contains the http client used to interact with Vault
type Client struct {
	HTTP *http.Client
}

// NewClient returns a http client configured according to the supplied Config, for use with Vault
func NewClient(certificate string) (*http.Client, error) {
	pool, err := cert.MakePool(certificate)
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

	client := &http.Client{
		Transport: transport,
	}

	return client, nil
}

// Do is a wrapper around http.do for the Vault http client that error checks and decodes to the destination interface
func (client Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := client.HTTP.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "while do-ing http request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("http error, status code %d", resp.StatusCode)
	}

	return resp, nil
}

// GetJSON exp
func (client Client) GetJSON(req *http.Request, dst interface{}) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	return nil
}

//TODO: rework and distill
func getSomething(resp *http.Response, v interface{}) error {
	// Return here if decoding the response isn't needed.
	if v == nil {
		return nil
	}

	// If v implements io.Writer, write the raw response body.
	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	// Get the value of v so we can test if it's a struct.
	dst := reflect.Indirect(reflect.ValueOf(v))

	// Return an error if v is not a struct or an io.Writer.
	if dst.Kind() != reflect.Struct {
		return fmt.Errorf("v must be a struct or an io.Writer")
	}

	// Try to get the Items and Pagination struct fields.
	items := dst.FieldByName("Items")
	pagination := dst.FieldByName("Pagination")

	// Unmarshal a single value if v does not contain the
	// Items and Pagination struct fields.
	if !items.IsValid() || !pagination.IsValid() {
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return readErr
		}
		// fmt.Println(body)
		return json.Unmarshal(body, v)
	}

	// Return an error if v.Items is not a slice.
	if items.Type().Kind() != reflect.Slice {
		return fmt.Errorf("v.Items must be a slice")
	}

	// Create a temporary buffer and copy all the read data into it.
	body := bytes.NewBuffer(nil)
	reader := io.TeeReader(resp.Body, body)

	// Unmarshal as a list of values as v.Items is a slice.
	raw, err := jsonapi.UnmarshalManyPayload(reader, items.Type().Elem())
	if err != nil {
		return err
	}

	// Make a new slice to hold the results.
	sliceType := reflect.SliceOf(items.Type().Elem())
	result := reflect.MakeSlice(sliceType, 0, len(raw))

	// Add all of the results to the new slice.
	for _, v := range raw {
		result = reflect.Append(result, reflect.ValueOf(v))
	}

	// Pointer-swap the result.
	items.Set(result)

	return nil
}

//TODO: consider usefulness
// checkResponseCode can be used to check the status code of an HTTP request.
func checkResponseCode(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode <= 299 {
		return nil
	}

	switch r.StatusCode {
	case 401:
		return errors.New("unauthorized")
	case 404:
		return errors.New("resource not found")
	}

	// Decode the error payload.
	errPayload := &jsonapi.ErrorsPayload{}
	err := json.NewDecoder(r.Body).Decode(errPayload)
	if err != nil || len(errPayload.Errors) == 0 {
		return fmt.Errorf(r.Status)
	}

	// Parse and format the errors.
	var errs []string
	for _, e := range errPayload.Errors {
		if e.Detail == "" {
			errs = append(errs, e.Title)
		} else {
			errs = append(errs, fmt.Sprintf("%s\n\n%s", e.Title, e.Detail))
		}
	}

	return fmt.Errorf(strings.Join(errs, "\n"))
}
