package assert

import (
	"strings"
	"testing"
)

// NoErr if you don't expect an error
func NoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, got: %v", err)
	}
}

// Err if you expect an error
func Err(t *testing.T, err error, errSubstr string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error, got none")
	} else if !strings.Contains(err.Error(), errSubstr) {
		t.Fatalf("Unexpected error, got: %v, wanted error containing substing: %s", err, errSubstr)
	}
}

// WantErr if switching between Err and NoErr with a bool is needed.
// Intended for use in table-based testing
func WantErr(t *testing.T, want bool, err error, errSubstr string) {
	t.Helper()
	if want {
		Err(t, err, errSubstr)
	} else {
		NoErr(t, err)
	}
}
