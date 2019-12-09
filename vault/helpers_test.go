package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_MakeUrl(t *testing.T) {
	// Test empty
	addr := ""
	path := ""
	want := "/v1/"
	got := makeURL(addr, path)
	assert.Result(t, got, want)

	// Test mock data
	addr = mock.Addr
	path = mock.Path
	want = addr + "/v1/" + path
	got = makeURL(addr, path)
	assert.Result(t, got, want)
}
