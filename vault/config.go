package vault

import (
	"os"

	"github.com/pkg/errors"
)

// Config contains the configuration information needed to do the initial setup of a Vault connection
type Config struct {
	Addr          string
	PemCert       string
	GithubToken   string
	K8ServicePath string
	K8MountPath   string
	K8Role        string
}

// NewConfig reads configuration information from provided file and returns a config struct containing this information.
func (vault *Vault) NewConfig() error {
	addr := os.Getenv("VAULT_ADDR")
	if addr == "" {
		addr = "https://127.0.0.1:8200"
	}

	pemCert := os.Getenv("VAULT_CACERT")
	githubToken := os.Getenv("GITHUB_TOKEN")

	if githubToken == "" {
		k8ServicePath := os.Getenv("SERVICE_ACCOUNT_PATH")
		if k8ServicePath == "" {
			k8ServicePath = "/var/run/secrets/kubernetes.io/serviceaccount/token" //TODO: verify valid default value
		}

		k8MountPath := os.Getenv("MOUNT_PATH")
		if k8MountPath == "" {
			k8MountPath = "kubernetes"
		}

		k8Role := os.Getenv("ROLE")
		if k8Role == "" {
			return errors.New("missing ROLE env var")
		}

		vault.K8ServicePath = k8ServicePath
		vault.K8MountPath = k8MountPath
		vault.K8Role = k8Role

	}

	vault.Addr = addr
	vault.PemCert = pemCert
	vault.GithubToken = githubToken

	return nil
}
