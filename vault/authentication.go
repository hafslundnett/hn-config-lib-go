package vault

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

//Auth explanation
type Auth struct {
	ClientToken string   `json:"client_token"`
	Accessor    string   `json:"accessor"`
	Policies    []string `json:"policies"`
}

//Metadata explanation
type Metadata struct {
	Role                     string `json:"role"`
	ServiceAccountName       string `json:"service_account_name"`
	ServiceAccountNamespace  string `json:"service_account_namespace"`
	ServiceAccountSecretName string `json:"service_account_secret_name"`
	ServiceAccountUID        string `json:"service_account_uid"`
}

//VaultToken explanation
type VaultToken struct {
	Auth          Auth     `json:"auth"`
	Metadata      Metadata `json:"metadata"`
	LeaseDuration int      `json:"lease_duration"`
	Renewable     bool     `json:"renewable"`
}

//Authenticate explanation
func Authenticate(cfg Config, client *http.Client) (VaultToken, error) {
	var token VaultToken

	url := makeURL(cfg, "auth/github/login")

	req, err := authReq(url, cfg.Login)
	if err != nil {
		return token, errors.Wrapf(err, "while building authentication request")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("while do-ing http request: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var b bytes.Buffer
		io.Copy(&b, resp.Body)
		return token, errors.Errorf("failed to get successful response: %#v, %s",
			resp, b.String())
	}

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		log.Println("failed to read body: ", err)
	}

	return token, nil
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

func makeURL(cfg Config, path string) string {
	version := "/v1/"
	address := cfg.VaultAddr

	return address + version + path
}
