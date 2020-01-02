package service

import (
	"errors"
	"os"

	"github.com/hafslundnett/hn-config-lib-go/libhttp"
)

// Config expl
type Config struct {
	Addr    string
	PemCert string
	Client  *libhttp.Client
}

// Configure expl
func Configure(service Service, env map[string]string) (err error) {
	config := new(Config)

	config.Addr = os.Getenv(env["addr"])
	if config.Addr == "" {
		err = errors.New("missing env var " + env["addr"])
		return
	}

	config.PemCert = os.Getenv(env["cert"])

	config.Client, err = libhttp.NewClient(config.PemCert)
	if err != nil {
		return
	}

	return service.Configure(config)
}
