package cert

import (
	"hafslundnett/x/hn-config-lib/testing/assert"
	"hafslundnett/x/hn-config-lib/testing/mock"
	"testing"
)

const (
	certFile = "_test_files/test_cert.cer"
	pemFile  = "_test_files/test_cert.pem"
)

func Test_MakePool(t *testing.T) {
	// Test with no file
	pool, err := MakePool()
	assert.NoErr(t, err)

	// Test with non-existing file
	pool, err = MakePool(mock.File)
	assert.Err(t, err, "failed to read CA file")

	// Test with existing file in wrong format
	pool, err = MakePool(certFile)
	assert.Err(t, err, "failed to parse PEM")

	// Test with valid, outdated, certificate from golang.org
	pool, err = MakePool(pemFile)
	assert.NoErr(t, err)

	if pool.Certs.Subjects()[0][0] != 48 {
		t.Fatalf("Incorrect information in pool")
	}
	if pool.Certs.Subjects()[0][1] != 109 {
		t.Fatalf("Incorrect information in pool")
	}
	if pool.Certs.Subjects()[0][2] != 49 {
		t.Fatalf("Incorrect information in pool")
	}
}
