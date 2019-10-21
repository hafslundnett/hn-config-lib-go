package main

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// RootCAs returns the list of trusted root CAs based off the provided configuration.
// If no CAs were specified, the system roots are used.
func RootCAs(vaultCaCert string) (*x509.CertPool, error) {
	certPool := x509.NewCertPool()
	if err := loadCertFile(certPool, vaultCaCert); err != nil {
		return nil, err
	}
	return certPool, nil
}

// loadCertFile loads the certificate at the given path into the given pool.
func loadCertFile(pool *x509.CertPool, certFile string) error {
	pem, err := ioutil.ReadFile(certFile)
	if err != nil {
		return errors.Wrapf(err, "failed to read CA file %s from disk", certFile)
	}

	if err := loadCert(pool, pem); err != nil {
		return errors.Wrapf(err, "failed to load CA at %s", certFile)
	}

	return nil
}

// loadCert loads a single pem-encoded certificate into the given pool.
func loadCert(pool *x509.CertPool, pem []byte) error {
	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return fmt.Errorf("failed to parse PEM")
	}
	return nil
}
