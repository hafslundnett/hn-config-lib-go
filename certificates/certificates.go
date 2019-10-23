package certificates

import (
	"crypto/x509"
	"io/ioutil"

	"github.com/pkg/errors"
)

//MakePool makes a new CA pool and appends provided certificates, if any.
//Accepts zero or more certificate files as arguments.
func MakePool(certFiles ...string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	for _, cert := range certFiles {
		if err := loadCertFile(pool, cert); err != nil {
			return nil, err
		}
	}

	return pool, nil
}

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

func loadCert(pool *x509.CertPool, pem []byte) error {
	if ok := pool.AppendCertsFromPEM(pem); !ok {
		return errors.Errorf("failed to parse PEM")
	}
	return nil
}
