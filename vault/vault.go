package vault

import "github.com/hafslundnett/hn-config-lib-go/libhttp"

// Vault contains all information needed to get and interact with Vault secrets, after initial configuration.
type Vault struct {
	Config
	Client *libhttp.Client
	Token  Token
}

// New initiaizes a new Vault prepares it for interacting with secrets.
// It reads configuration information from the environment, configures a HTTP client and gets an authentification token to get secrets.
func New() (Vault, error) {
	service, err := Setup(Vault{})
	return service.(Vault), err
}

// MakeClient returns a http client configured according to the supplied Config, for use with Vault
func (vault Vault) MakeClient() (err error) {
	vault.Client, err = libhttp.NewClient(vault.Config.PemCert)
	return
}
