package cert

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// AppendFromFiles loads the certs from a slice of filenames to the cert pool
func (pool Pool) AppendFromFiles(certFiles []string) error {
	for _, cert := range certFiles {
		if cert != "" {
			if err := pool.loadFromFile(cert); err != nil {
				return err
			}
		}
	}

	return nil
}

// loadFromFile loads the certs from a specified file to the provided pool
func (pool Pool) loadFromFile(cert string) error {
	pem, err := ioutil.ReadFile(cert)
	if err != nil {
		return errors.Wrapf(err, "failed to read CA file \"%s\" from disk", cert)
	}

	if err := pool.loadFromPEM(pem); err != nil {
		return errors.Wrapf(err, "failed to load CA at \"%s\"", cert)
	}

	return nil
}

// loadCert appends the PEM-formatted certificate to the provided pool
func (pool Pool) loadFromPEM(pem []byte) error {
	if ok := pool.Certs.AppendCertsFromPEM(pem); !ok {
		return errors.Errorf("failed to parse PEM")
	}

	return nil
}
