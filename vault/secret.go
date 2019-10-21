package vault

import (
	"encoding/json"
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
	var secret Secret

	url := makeURL(vault.Cfg, path)

	req, err := secretsReq(url, vault.Token.Auth.ClientToken)
	if err != nil {
		log.Println("while logging in: ", err)
	}

	resp, err := vault.Client.Do(req)
	if err != nil {
		log.Println("while logging in: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("error, status code: ", resp.StatusCode)
		//var b bytes.Buffer
		//io.Copy(&b, resp.Body)
		//return secret, fmt.Errorf("failed to get successful response: %#v, %s",
		//	resp, b.String())
	}

	if err := json.NewDecoder(resp.Body).Decode(&secret); err != nil {
		log.Println("failed to read body: ", err)
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
