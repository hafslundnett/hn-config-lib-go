package vault

import (
	"bytes"
	"encoding/json"
	"testing"
)

func Test_gitLogin(t *testing.T) {

}
func Test_k8Login(t *testing.T) {

}

func Test_loginBuffer(t *testing.T) {
	//Test with empty arg
	got, err := loginBuffer(nil)
	assertNoErr(t, err)
	if got.Len() != 4 {
		t.Fatalf("result is %d bytes, should be 4 bytes", got.Len())
	}

	//Test contents of result
	js, err := json.Marshal(mockToken)
	assertNoErr(t, err)
	want := bytes.NewBuffer(js)
	got, err = loginBuffer(mockToken)
	assertNoErr(t, err)
	if got.String() != want.String() {
		t.Fatal("got incorrect bytestring")
	}
}
