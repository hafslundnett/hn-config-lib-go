package vault

import (
	"net/http"

	"github.com/pkg/errors"
)

//Token is used for authenticating Vault requests
type Token struct {
	Auth          Auth     `json:"auth"`
	Metadata      Metadata `json:"metadata"`
	LeaseDuration int      `json:"lease_duration"`
	Renewable     bool     `json:"renewable"`
}

//Auth contains the token information for authenticating Vault requests
type Auth struct {
	ClientToken string   `json:"client_token"`
	Accessor    string   `json:"accessor"`
	Policies    []string `json:"policies"`
}

//Metadata contains important metadata for the Vault Token
type Metadata struct {
	Role                     string `json:"role"`
	ServiceAccountName       string `json:"service_account_name"`
	ServiceAccountNamespace  string `json:"service_account_namespace"`
	ServiceAccountSecretName string `json:"service_account_secret_name"`
	ServiceAccountUID        string `json:"service_account_uid"`
}

//Authenticate uses supplied login information to authenticate to Vault and get an authentification token
func (vault *Vault) Authenticate() error {
	path := "auth/github/login"
	url := makeURL(vault.Cfg, path)

	req, err := authReq(url, vault.Cfg.Login)
	if err != nil {
		return err
	}

	if err := vault.Client.remoteCall(req, &vault.Token); err != nil {
		return err
	}

	return nil
}

func authReq(url string, login string) (*http.Request, error) {
	body, err := loginBuffer(login)
	if err != nil {
		return nil, errors.Wrapf(err, "while converting token to buffer")
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, errors.Wrapf(err, "while building http request")
	}

	return req, nil
}
