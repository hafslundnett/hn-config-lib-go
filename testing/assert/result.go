package assert

import (
	"reflect"
	"testing"
)

// Result compares test result with expected result using normal != assesment
func Result(t *testing.T, got, want interface{}) {
	t.Helper()
	if want != got {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}

// DeepResult compares test result with expected result using deepEqual
func DeepResult(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %s, got %s", want, got)
	}
}
