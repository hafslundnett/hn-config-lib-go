package cert

import (
	"crypto/x509"

	"github.com/pkg/errors"
)

// Pool is a storage for RootCAs, used in http requests
type Pool struct {
	Certs *x509.CertPool
}

// MakePool copies the system CA pool and appends provided certificates, if any.
// Accepts zero or more certificate files as arguments.
// Empty strings are ignored
func MakePool(certFiles ...string) (*Pool, error) {
	pool := new(Pool)
	var err error

	pool.Certs, err = x509.SystemCertPool() // Returns pool as nil and an error on Windows
	if err != nil {
		pool.Certs = x509.NewCertPool() // Use an empty one instead. No OS provided certs from windows
	}

	err = pool.AppendCAs()
	if err != nil {
		return nil, err
	}

	err = pool.AppendFromFiles(certFiles)
	if err != nil {
		return nil, errors.Wrap(err, "while loading CA cert from file")
	}

	return pool, nil
}
