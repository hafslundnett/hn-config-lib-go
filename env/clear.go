package env

import (
	"os"

	"github.com/pkg/errors"
)

var original = map[string]string{}

// Save takes the contents of the provided env variables and saves them in a map.
// Intended to use before env.Clear.
// Restore with env.Reset
func Save(vars ...string) error {
	for _, k := range vars {
		if k == "" {
			return errors.New("while saving \"\": invalid; throws error on os.setenv")
		}
		original[k] = os.Getenv(k)
	}
	return nil
}

// Clear sets provided environment variables to the empty string "".
func Clear(vars ...string) (err error) {
	for _, k := range vars {
		err = Set(k, "")
	}
	return
}

// Reset takes the env variables saved by env.Save and restores them to the environment.
// If env.Save has not previously been called, this does nothing.
func Reset() error {
	return SetMap(original)
}
