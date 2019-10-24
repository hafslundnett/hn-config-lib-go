package vault

import (
	"os"

	"github.com/pkg/errors"
)

//Config contains the configuration information needed to do the initial setup of a Vault connection
type Config struct {
	VaultAddr string
	PemCert   string
	GitToken  string
	K8Token   string
	Role      string
}

//NewConfig reads configuration information from provided file and returns a config struct containing this information.
func (vault *Vault) NewConfig() error {
	vault.Config.VaultAddr = os.Getenv("VAULT_ADDR")
	if vault.Config.VaultAddr == "" {
		vault.Config.VaultAddr = "https://127.0.0.1:8200"
	}

	vault.Config.PemCert = os.Getenv("VAULT_CACERT")

	vault.Config.GitToken = os.Getenv("GITHUB_TOKEN")

	if vault.Config.GitToken == "" {
		saPath := os.Getenv("SERVICE_ACCOUNT_PATH")
		if saPath == "" {
			saPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		}

		vaultK8SMountPath := os.Getenv("MOUNT_PATH")
		if vaultK8SMountPath == "" {
			errors.New("missing MOUNT_PATH")
		}

		role := os.Getenv("ROLE")
		if role == "" {
			errors.New("missing ROLE")
		}
	}

	return nil
}
