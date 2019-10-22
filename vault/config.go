package vault

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//Config contains the configuration information needed to do the initial setup of a Vault connection
type Config struct {
	Login     string `yaml:"login_token"`
	VaultAddr string `yaml:"vault_addr"`
	CaCert    string `yaml:"cert"`
}

//NewConfig reads configuration information from provided file and returns a config struct containing this information.
func (vault *Vault) NewConfig(configFile string) error {
	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.Wrap(err, "while reading configuration file")
	}

	if err = yaml.Unmarshal(bs, &vault.Cfg); err != nil {
		return errors.Wrap(err, "while parsing configuration file")
	}

	return nil
}
