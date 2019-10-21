package main

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type config struct {
	LoginToken string `yaml:"token"`
	VaultAddr  string `yaml:"vault_addr"`
	CaCert     string `yaml:"cert"`
}

func getConfig(configFile string) (config, error) {
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config{}, errors.Wrap(err, "while reading configuration file")
	}

	var cfg config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return config{}, errors.Wrap(err, "while to parsing configuration file")
	}

	return cfg, nil
}
