package vault

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Token is used for authenticating Vault requests
type Token struct {
	Auth          Auth     `json:"auth"`
	Metadata      Metadata `json:"metadata"`
	LeaseDuration int      `json:"lease_duration"`
	Renewable     bool     `json:"renewable"`
}

// Auth contains the token information for authenticating Vault requests
type Auth struct {
	ClientToken string   `json:"client_token"`
	Accessor    string   `json:"accessor"`
	Policies    []string `json:"policies"`
}

// Metadata contains important metadata for the Vault Token
type Metadata struct {
	Role                     string `json:"role"`
	ServiceAccountName       string `json:"service_account_name"`
	ServiceAccountNamespace  string `json:"service_account_namespace"`
	ServiceAccountSecretName string `json:"service_account_secret_name"`
	ServiceAccountUID        string `json:"service_account_uid"`
}

// Authenticate uses supplied login information to authenticate to Vault and get an authentification token
func (vault *Vault) Authenticate() error {
	req, err := authReq(vault.Config)
	if err != nil {
		return err
	}

	if err := vault.Client.Do(req, &vault.Token); err != nil {
		return err
	}

	return nil
}

// authReq reurns a http request for authenticating to Vault
func authReq(cfg Config) (*http.Request, error) {
	body, path, err := authBody(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "while converting token to buffer")
	}

	url := makeURL(cfg.Addr, path)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "while building http request")
	}

	return req, nil
}

// authBody reads the configuration and returns a reader to the correct kind of login
func authBody(cfg Config) (io.Reader, string, error) {
	if cfg.GithubToken != "" {
		path := "auth/github/login"
		body, err := githubLogin(cfg.GithubToken)
		return body, path, err
	}

	path := "auth/" + cfg.K8MountPath + "/login"
	body, err := k8Login(cfg.K8ServicePath, cfg.K8Role)
	return body, path, err
}
