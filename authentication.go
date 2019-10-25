package vault

import (
	"io"
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
	req, err := authReq(vault.Config)
	if err != nil {
		return err
	}

	if err := vault.Client.remoteCall(req, &vault.Token); err != nil {
		return err
	}

	return nil
}

func authReq(cfg Config) (*http.Request, error) {
	body, path, err := makeBody(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "while converting token to buffer")
	}

	url := makeURL(cfg, path)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "while building http request")
	}

	return req, nil
}

func makeBody(cfg Config) (io.Reader, string, error) {
	var body io.Reader
	var path string
	var err error

	if cfg.GitHubToken != "" {
		path = "auth/github/login"
		body, err = gitLogin(cfg.GitHubToken)
	} else {
		path = "auth/kubernetes/login" //TODO: construct proper path
		body, err = k8Login(cfg.K8Token, cfg.Role)
	}

	return body, path, err
}
