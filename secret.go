package vault

import (
	"net/http"

	"github.com/pkg/errors"
)

//Secret contains all data from a Vault secret
type Secret struct {
	RequestID     string            `json:"client_token"`
	LeaseID       string            `json:"lease_id"`
	Renewable     bool              `json:"renewable"`
	LeaseDuration int               `json:"lease_duration"`
	Data          map[string]string `json:"data"`
}

//GetSecret returns the secret from the provided path
func (vault Vault) GetSecret(path string) (*Secret, error) {
	secret := new(Secret)
	url := makeURL(vault.Config, path)

	req, err := secretsReq(url, vault.Token.Auth.ClientToken)
	if err != nil {
		return nil, err
	}

	if err = vault.Client.do(req, &secret); err != nil {
		return nil, err
	}

	return secret, nil
}

func secretsReq(url string, auth string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "while building http request")
	}

	req.Header.Set("X-Vault-Token", auth)
	if err != nil {
		return nil, errors.Wrap(err, "while setting header")
	}

	return req, nil
}
