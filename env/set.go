package env

import (
	"os"

	"github.com/pkg/errors"
)

// Set enables setting multiple environment variables by putting them after one another.
// Example: env.FromStrings("USER", "guest", "PASSWORD", "qwer1234")
func Set(vars ...string) error {
	for i := 0; i+1 < len(vars); i += 2 {
		err := os.Setenv(vars[i], vars[i+1])
		if err != nil {
			return errors.Wrapf(err, "while setting env var %s", vars[i])
		}
	}

	return nil
}

// SetMap enables setting multiple environment variables from a map.
// Example: m := map[string]string{"USER": "guest", "PASSWORD": "qwer1234"}
// env.FromMap(m)
func SetMap(vars map[string]string) error {
	for k, v := range vars {
		err := os.Setenv(k, v)
		if err != nil {
			return errors.Wrapf(err, "while setting env var %s", k)
		}
	}

	return nil
}
