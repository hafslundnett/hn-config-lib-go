package vault

import (
	"os"

	"github.com/pkg/errors"
)

var envars = map[string]string{
	"addr":     "VAULT_ADDR",
	"cert":     "VAULT_CACERT",
	"ghToken":  "GITHUB_TOKEN",
	"servPath": "SERVICE_ACCOUNT_PATH",
	"mntPath":  "MOUNT_PATH",
	"role":     "ROLE",
}

// Config contains the configuration information needed to do the initial setup of a Vault connection
type Config struct {
	*Conf
	GithubToken   string
	K8ServicePath string
	K8MountPath   string
	K8Role        string
}

// Configure reads configuration information from provided file and returns a config struct containing this information.
func (vault Vault) Configure(c *Conf) error {
	vault.Conf = c

	vault.GithubToken = os.Getenv(envars["ghToken"])
	if vault.GithubToken == "" {

		k8ServicePath := os.Getenv(envars["servPath"])
		if k8ServicePath == "" {
			k8ServicePath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		}

		k8MountPath := os.Getenv(envars["mntPath"])
		if k8MountPath == "" {
			k8MountPath = "kubernetes"
		}

		k8Role := os.Getenv(envars["role"])
		if k8Role == "" {
			return errors.New("missing either ROLE or GITHUB_TOKEN env var")
		}

		vault.K8ServicePath = k8ServicePath
		vault.K8MountPath = k8MountPath
		vault.K8Role = k8Role

	}

	return nil
}
