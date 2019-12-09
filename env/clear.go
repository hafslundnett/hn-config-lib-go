package env

import "os"

var original = map[string]string{}

// Save expl
func Save(vars ...string) {
	for _, v := range vars {
		original[v] = os.Getenv(v)
	}
}

// Clear expl
func Clear(vars ...string) {
	for _, v := range vars {
		Set(v, "")
	}
}

// Reset expl
func Reset() {
	SetMap(original)
}
