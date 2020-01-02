package vault

import (
	"os"

	"github.com/hafslundnett/hn-config-lib-go/libhttp"
	"github.com/pkg/errors"
)

/**
// Service expl
type Service interface {
	Configure() error
	MakeClient() error
	GetAccess() error
}

// Setup expl
func Setup(service Service) (Service, error) {
	if err := service.Configure(); err != nil {
		return service, err
	}

	if err := service.MakeClient(); err != nil {
		return service, err
	}

	if err := service.GetAccess(); err != nil {
		return service, err
	}

	return service, nil
}
*/

// Service expl
type Service interface {
	Configure(*Conf) error
	GetEnvars() map[string]string
	GetAccess() error
}

// Setup expl
func Setup(service Service) (Service, error) {
	Configure(service, service.GetEnvars())

	err := service.GetAccess()

	return service, err
}

// Config expl
type Conf struct {
	Addr    string
	PemCert string
	Client  *libhttp.Client
}

func Configure(service Service, env map[string]string) (err error) {
	config := new(Conf)

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

func (Vault) GetEnvars() map[string]string {
	return envars
}
