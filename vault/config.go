package vault

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//Config contains the configuration information needed to do the initial setup of a Vault connection
type Config struct {
	VaultAddr string `yaml:"vault_addr"`
	PemCert   string `yaml:"pem_cert"`
	GitToken  string `yaml:"git_token"`
	K8Token   string `yaml:"k8_token"`
	Role      string `yaml:"k8_role"`
}

//NewConfig reads configuration information from provided file and returns a config struct containing this information.
func (vault *Vault) NewConfig(configFile string) error {
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.Wrap(err, "while reading configuration file")
	}

	if err = yaml.Unmarshal(bs, &vault.Config); err != nil {
		return errors.Wrap(err, "while parsing configuration file")
	}

	return nil
}
