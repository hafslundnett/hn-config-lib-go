package env

import "os"

var original = map[string]string{}

// Save takes the contents of the provided env variables and saves them in a map.
// Intended to use before env.Clear.
// Restore with env.Reset
func Save(vars ...string) {
	for _, v := range vars {
		original[v] = os.Getenv(v)
	}
}

// Clear sets provided environment variables to the empty string "".
func Clear(vars ...string) {
	for _, v := range vars {
		Set(v, "")
	}
}

// Reset takes the env variables saved by env.Save and restores them to the environment.
// If env.Save has not previously been called, this does nothing.
func Reset() {
	SetMap(original)
}
