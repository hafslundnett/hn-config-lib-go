package vault

import "hafslundnett/x/hn-config-lib/http"

// Client contains the http client used to interact with Vault
type Client struct {
	http.Client
}

// NewClient returns a http client configured according to the supplied Config, for use with Vault
func (vault *Vault) NewClient() error {
	var err error
	vault.Client.HTTP, err = http.NewClient(vault.Config.PemCert)
	return err
}
