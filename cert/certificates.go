package cert

import (
	"crypto/x509"
	"log"
	"runtime"

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

	if runtime.GOOS == "windows" && len(certFiles) == 0 {
		log.Println("Windows requires certificate file to avoid error \"x509: certificate signed by unknown authority\"")
	}

	err := pool.AppendFromFiles(certFiles)
	if err != nil {
		return nil, errors.Wrap(err, "while loading CA cert from file")
	}

	return pool, nil
}
