package vault

import "github.com/hafslundnett/hn-config-lib-go/hnhttp"

// Vault contains all information needed to get and interact with Vault secrets, after initial configuration.
type Vault struct {
	Config
	Client *hnhttp.Client
	Token  Token
}

// New initiaizes a new Vault prepares it for interacting with secrets.
// It reads configuration information from the environment, configures a HTTP client and gets an authentification token to get secrets.
func New() (*Vault, error) {
	vault := new(Vault)

	if err := vault.NewConfig(); err != nil {
		return vault, err
	}

	if err := vault.NewClient(); err != nil {
		return vault, err
	}

	if err := vault.Authenticate(); err != nil {
		return vault, err
	}

	return vault, nil
}

// NewClient returns a http client configured according to the supplied Config, for use with Vault
func (vault *Vault) NewClient() error {
	var err error
	vault.Client, err = hnhttp.NewClient(vault.Config.PemCert)
	return err
}
