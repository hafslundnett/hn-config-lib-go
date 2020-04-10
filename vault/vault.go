package vault

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/hafslundnett/hn-config-lib-go/hnhttp"
	"io/ioutil"
	"os"
)

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

// SetDefaultGoogleCredentials assumes that the given vault path is to a base 64 encoded
// json credentials structure. This information is downloaded and saved to a file. A reference
// to this file is stored in the environment variable GOOGLE_APPLICATION_CREDENTIALS
func (vault *Vault) SetDefaultGoogleCredentials(path, key string) error {
	secret, err := vault.GetSecret(path)
	if err != nil {
		return err
	}
	var encoded string
	var ok bool
	if encoded, ok = secret.Data[key]; !ok {
		return errors.New(fmt.Sprintf("no key %s found in secret at %s", key, path))
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	fn := getRandomFileName("google_creds", ".json")
	err = ioutil.WriteFile(fn, []byte(decoded), 0644)
	if err != nil {
		return err
	}
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", fn)
	return nil
}