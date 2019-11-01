package env

import (
	"os"
)

// Set enables setting multiple environment variables by putting the after one another.
// Example: env.FromStrings("USER", "guest", "PASSWORD", "qwer1234")
func Set(vars ...string) {
	for i := 0; i < len(vars); i += 2 {
		os.Setenv(vars[i], vars[i+1])
	}
}

// SetMap enables setting multiple environment variables from a map
// Example: m := map[string]string{"USER": "guest", "PASSWORD": "qwer1234"}
// env.FromMap(m)
func SetMap(vars map[string]string) {
	for k, v := range vars {
		os.Setenv(k, v)
	}
}
