package vault

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Token is used for authenticating Vault requests
type Token struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	LeaseDuration int    `json:"lease_duration"`
	Renewable     bool   `json:"renewable"`
	Auth          Auth   `json:"auth"`
}

// Auth contains the token information for authenticating Vault requests
type Auth struct {
	ClientToken   string                 `json:"client_token"`
	TokenType     string                 `json:"token_type"`
	Accessor      string                 `json:"accessor"`
	EntityID      string                 `json:"entity_id"`
	LeaseDuration int                    `json:"lease_duration"`
	Renewable     bool                   `json:"renewable"`
	Orphan        bool                   `json:"orphan"`
	Policies      []string               `json:"policies"`
	TokenPolicies []string               `json:"token_policies"`
	Metadata      map[string]interface{} `json:"metadata"`
}

// Authenticate uses supplied login information to authenticate to Vault and get an authentification token
func (vault *Vault) Authenticate() error {
	req, err := authReq(vault.Config)
	if err != nil {
		return err
	}

	if err := vault.Client.Do(req, &vault.Token); err != nil {
		return errors.Wrapf(err, "while authenticating to Vault at %s", req.URL)
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

// makeURL returns a correctly formatted url for Vault http requests
func makeURL(address, path string) string {
	version := "/v1/"
	return address + version + path
}
