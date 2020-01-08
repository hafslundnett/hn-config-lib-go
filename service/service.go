package service

import (
	"github.com/hafslundnett/hn-config-lib-go/libhttp"
)

// Service expl
type Service interface {
	Configure(libhttp.Client) error
	ConnectToServer() error
}

// Setup expl
func Setup(service Service, pemCert string) error {
	client, err := libhttp.NewClient(pemCert)
	if err != nil {
		return err
	}

	err = service.Configure(client)
	if err != nil {
		return err
	}

	err = service.ConnectToServer()
	if err != nil {
		return err
	}

	return nil
}
