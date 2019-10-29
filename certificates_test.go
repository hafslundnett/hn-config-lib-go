package vault

import (
	"testing"
)

func Test_MakePool(t *testing.T) {
	SetEnv("", "", "", "", "", "")

	// Test with no file
	pool, err := MakePool()
	assertNoErr(t, err)

	// Test with non-existing file
	pool, err = MakePool(mockFile)
	assertErr(t, err, "failed to read CA file")

	// Test with existing file in wrong format
	pool, err = MakePool("test_files/test_cert.cer")
	assertErr(t, err, "failed to parse PEM")

	// Test with valid, outdated, certificate from golang.org
	pool, err = MakePool("test_files/test_cert.pem")
	assertNoErr(t, err)

	if pool.Subjects()[0][0] != 48 {
		t.Fatalf("Incorrect information in pool")
	}
	if pool.Subjects()[0][1] != 109 {
		t.Fatalf("Incorrect information in pool")
	}
	if pool.Subjects()[0][2] != 49 {
		t.Fatalf("Incorrect information in pool")
	}
}
