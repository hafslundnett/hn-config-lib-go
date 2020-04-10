package vault

import (
	"os"

	"github.com/hafslundnett/hn-config-lib-go/service"
)

// SecretsManager represents a service that is able to provide clients with a secret stored at a privded path.
type SecretsManager interface {
	GetSecret(path string) (*Secret, error)
	SetDefaultGoogleCredentials(path, key string) error
}

// Vault contains all information needed to get and interact with Vault secrets, after initial configuration.
type Vault struct {
	Config
	Token Token
}

// New initiaizes a new Vault prepares it for interacting with secrets.
// It reads configuration information from the environment, configures a HTTP client and gets an authentification token to get secrets.
func New() (SecretsManager, error) {
	vault := new(Vault)
	cert := os.Getenv(envars["cert"])
	err := service.Setup(vault, cert)
	return vault, err
}

// ConnectToServer performs neccessary setup for connections to the external HID service
func (vault *Vault) ConnectToServer() error {
	return vault.Authenticate()
}
