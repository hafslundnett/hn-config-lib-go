package vault

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

// Client contains the http client used to interact with Vault
type Client struct {
	HTTP *http.Client
}

// NewClient returns a http client configured according to the supplied Config, for use with Vault
func (vault *Vault) NewClient() error {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	if vault.Config.PemCert != "" {
		pool, err := MakePool(vault.Config.PemCert)
		if err != nil {
			return errors.Wrap(err, "while getting CA Certs")
		}
		tlsConfig.RootCAs = pool
	}

	transport := &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	vault.Client.HTTP = &http.Client{
		Transport: transport,
	}

	return nil
}

// do is a wrapper around http.do for the Vault http client that error checks and decodes to the destination interface
func (client Client) do(req *http.Request, dst interface{}) error {
	resp, err := client.HTTP.Do(req)
	if err != nil {
		return errors.Wrap(err, "while do-ing http request")
	}

	if resp.StatusCode != 200 {
		return errors.Errorf("http error, status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	return nil
}
