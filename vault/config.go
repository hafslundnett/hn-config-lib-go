package vault

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

//Config explanation
type Config struct {
	Login     string `yaml:"login_token"`
	VaultAddr string `yaml:"vault_addr"`
	CaCert    string `yaml:"cert"`
}

//NewConfig explanation
func NewConfig(configFile string) (Config, error) {
	cfg := Config{}

	bs, err := ioutil.ReadFile(configFile)
	if err != nil {
		return cfg, errors.Wrap(err, "while reading configuration file")
	}

	err = yaml.Unmarshal(bs, &cfg)
	if err != nil {
		return cfg, errors.Wrap(err, "while to parsing configuration file")
	}

	return cfg, nil
}
