package assert

import "testing"

// Result compares test result with expected result
func Result(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}
