package vault

import (
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
)

// MakePool makes a new CA pool and appends provided certificates, if any.
// Accepts zero or more certificate files as arguments.
// Empty strings are ignored
func MakePool(certFiles ...string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	for _, cert := range certFiles {
		if cert != "" {
			if err := loadCertFile(pool, cert); err != nil {
				return nil, err
			}
		}
	}

	return pool, nil
}

// loadCertFile loads the certs from a specified file to the provided pol
func loadCertFile(pool *x509.CertPool, cert string) error {
	pem, err := ioutil.ReadFile(cert)
	if err != nil {
		return errors.Wrapf(err, "failed to read CA file \"%s\" from disk", cert)
	}

	if err := loadCert(pool, pem); err != nil {
		return errors.Wrapf(err, "failed to load CA at \"%s\"", cert)
	}

	return nil
}

// loadCert appends the PEM-formatted certificate to the provided pool
func loadCert(pool *x509.CertPool, pem []byte) error {
	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return errors.Errorf("failed to parse PEM")
	}
	return nil
}
