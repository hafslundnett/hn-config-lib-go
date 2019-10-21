package vault

import (
	"net/http"

	"github.com/pkg/errors"
)

//Vault explanation
type Vault struct {
	Cfg    Config
	Client *http.Client
	Token  VaultToken
}

//NewVault explanation
func NewVault(cfgFile string) (*Vault, error) {
	cfg, err := NewConfig(cfgFile)
	if err != nil {
		return nil, errors.Wrapf(err, "while getting configuration: ")
	}

	client, err := MakeClient(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "while making client")
	}

	token, err := Authenticate(cfg, client)
	if err != nil {
		return nil, errors.Wrapf(err, "while : ")
	}

	vault := &Vault{
		Cfg:    cfg,
		Client: client,
		Token:  token,
	}

	return vault, nil
}
