package vault

import (
	"os"

	"github.com/hafslundnett/hn-config-lib-go/libhttp"

	"github.com/pkg/errors"
)

var envars = map[string]string{
	"addr":    "VAULT_ADDR",
	"github":  "GITHUB_TOKEN",
	"account": "SERVICE_ACCOUNT_PATH",
	"mount":   "MOUNT_PATH",
	"role":    "ROLE",
}

// Config contains the configuration information needed to do the initial setup and renewal of a Vault service
type Config struct {
	Addr          string
	GithubToken   string
	K8ServicePath string
	K8MountPath   string
	K8Role        string

	Client libhttp.Client
}

// Configure reads configuration information from provided file and returns a config struct containing this information.
func (vault *Vault) Configure(client libhttp.Client) error {
	addr := os.Getenv(envars["addr"])
	if addr == "" {
		return errors.New("missing env var " + envars["addr"])
	}

	githubToken := os.Getenv(envars["github"])
	if githubToken == "" {

		k8ServicePath := os.Getenv(envars["account"])
		if k8ServicePath == "" {
			k8ServicePath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		}

		k8MountPath := os.Getenv(envars["mount"])
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

	vault.Addr = addr
	vault.GithubToken = githubToken
	vault.Client = client

	return nil
}
