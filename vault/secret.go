package vault

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
)

//Secret explanation
type Secret struct {
	RequestID     string            `json:"client_token"`
	LeaseID       string            `json:"lease_id"`
	Renewable     bool              `json:"renewable"`
	LeaseDuration int               `json:"lease_duration"`
	Data          map[string]string `json:"data"`
}

//GetSecret explanation
func (vault Vault) GetSecret(path string) (Secret, error) {
	secret := Secret{}
	url := makeURL(vault.Cfg, path)

	req, err := secretsReq(url, vault.Token.Auth.ClientToken)
	if err != nil {
		log.Println("while logging in: ", err)
	}

	if err = vault.Client.remoteCall(req, &secret); err != nil {
		return secret, err
	}

	return secret, nil
}

func secretsReq(url string, auth string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "while building http request")
	}

	req.Header.Set("X-Vault-Token", auth)
	if err != nil {
		return nil, errors.Wrapf(err, "while setting header")
	}

	return req, nil
}
