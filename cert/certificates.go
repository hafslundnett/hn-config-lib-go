package cert

import (
	"crypto/x509"

	"github.com/pkg/errors"
)

// Pool expl
type Pool struct {
	Certs *x509.CertPool
}

// MakePool makes a new CA pool and appends provided certificates, if any.
// Accepts zero or more certificate files as arguments.
// Empty strings are ignored
func MakePool(certFiles ...string) (*Pool, error) {
	pool := &Pool{
		Certs: x509.NewCertPool(),
	}

	err := pool.AppendFromSystem()
	if err != nil {
		return nil, errors.Wrap(err, "while loading system CA certs")
	}

	err = pool.AppendFromFiles(certFiles)
	if err != nil {
		return nil, errors.Wrap(err, "while loading CA cert from file")
	}

	return pool, nil
}
