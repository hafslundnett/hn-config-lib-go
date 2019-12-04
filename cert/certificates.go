package cert

import (
	"crypto/x509"
	"log"

	"github.com/pkg/errors"
)

// Pool expl
type Pool struct {
	Certs *x509.CertPool
}

// MakePool copies the system CA pool and appends provided certificates, if any.
// Accepts zero or more certificate files as arguments.
// Empty strings are ignored
// NOTE: Cannot get certificates on Windows, therefore a certificate file is required!
func MakePool(certFiles ...string) (*Pool, error) {
	pool := new(Pool)
	var err error

	pool.Certs, err = x509.SystemCertPool() // Throws error on Windows
	if err != nil {
		pool.Certs = x509.NewCertPool()

		if len(certFiles) == 0 {
			log.Println("Windows can not load system certificates, your pool is now empty.")
			log.Println("HTTPS requests will respond with \"x509: certificate signed by unknown authority\".")
			log.Println("Consider not setting TLS' RootCAs to make it use the default certificates instead.")
		}
	}

	err = pool.AppendFromFiles(certFiles)
	if err != nil {
		return nil, errors.Wrap(err, "while loading CA cert from file")
	}

	return pool, nil
}
