package vault

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
)

//Client contains the http client used to interact with Vault
type Client struct {
	HTTP *http.Client
}

//NewClient returns a http client configured according to the supplied Config, for use with Vault
func (vault *Vault) NewClient() error {
	pool, err := MakePool(vault.Config.PemCert)
	if err != nil {
		return errors.Wrap(err, "while getting CA Certs")
	}

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		RootCAs:    pool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	if err = http2.ConfigureTransport(transport); err != nil {
		return errors.Wrap(err, "while configuring http2")
	}

	vault.Client.HTTP = &http.Client{
		Transport: transport,
	}

	return nil
}

func (client Client) remoteCall(req *http.Request, dst interface{}) error {
	resp, err := client.HTTP.Do(req)
	if err != nil {
		return errors.Wrap(err, "while do-ing http request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var b bytes.Buffer
		io.Copy(&b, resp.Body)
		return errors.Errorf("error %d, failed to get successful response: %#v, %s", resp.StatusCode, resp, b.String())
	}

	if err := json.NewDecoder(resp.Body).Decode(&dst); err != nil {
		return errors.Wrap(err, "failed to read body")
	}

	return nil
}

func makeURL(cfg Config, path string) string {
	version := "/v1/"
	address := cfg.VaultAddr

	return address + version + path
}
