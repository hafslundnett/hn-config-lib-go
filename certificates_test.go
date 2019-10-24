package vault

import "testing"

func Test_MakePool(t *testing.T) {
	//Test with no file
	pool, err := MakePool()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	//Test with non-existing file
	pool, err = MakePool("test_files/not_a_file.tull")
	if err.Error() != "failed to read CA file \"test_files/not_a_file.tull\" from disk: open test_files/not_a_file.tull: Systemet finner ikke angitt fil." {
		t.Errorf("Unexpected error: %v", err)
	}

	//Test with existing file in wrong format
	pool, err = MakePool("test_files/test_cert.cer")
	if err.Error() != "failed to load CA at \"test_files/test_cert.cer\": failed to parse PEM" {
		t.Errorf("Unexpected error: %v", err)
	}

	//Test with valid, outdated, certificate from golang.org
	pool, err = MakePool("test_files/test_cert.pem")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if pool.Subjects()[0][0] != 48 {
		t.Errorf("Incorrect information in pool")
	}
	if pool.Subjects()[0][1] != 109 {
		t.Errorf("Incorrect information in pool")
	}
	if pool.Subjects()[0][2] != 49 {
		t.Errorf("Incorrect information in pool")
	}
}
