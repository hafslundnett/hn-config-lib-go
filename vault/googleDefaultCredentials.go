package vault

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type fileTextSaver func(path string, contents []byte) error

func (vault *Vault) SetDefaultGoogleCredentials(path, key string) error {
	saver := func(path string, contents []byte) error {
		return ioutil.WriteFile(path, contents, 0644)
	}
	return vault.setDefaultGoogleCreds(path, key, saver)
}

func (vault *Vault) setDefaultGoogleCreds(path, key string, saver fileTextSaver) error {
	secret, err := vault.GetSecret(path)
	if err != nil {
		return err
	}
	var encoded string
	var ok bool
	if encoded, ok = secret.Data[key]; !ok {
		return errors.New(fmt.Sprintf("no key %s found in secret at %s", key, path))
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	fn := getRandomFileName("google_creds", "json")
	err = saver(fn, decoded)
	if err != nil {
		return err
	}
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", fn)
	return nil
}

func getRandomFileName(prefix, suffix string) string {
	i := rand.Intn(100000)
	return fmt.Sprintf("%s_%d.%s", prefix, i, suffix)
}