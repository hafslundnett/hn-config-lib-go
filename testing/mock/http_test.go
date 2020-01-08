package mock

import (
	"net/http"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

func TestRequest(t *testing.T) {
	type args struct {
		method string
		url    string
		body   string
		header []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "additional header entries",
			args: args{"GET", URL, Secret, []string{Role, ID, File}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := http.NewRequest(tt.args.method, tt.args.url, Body(tt.args.body))
			assert.NoErr(t, err)

			want.Header.Set(tt.args.header[0], tt.args.header[1])

			got := Request(t, tt.args.method, tt.args.url, tt.args.body, tt.args.header...)

			assert.Result(t, got.Method, want.Method)
			assert.DeepResult(t, got.URL, want.URL)
			assert.DeepResult(t, got.Body, want.Body)
			assert.DeepResult(t, got.Header, want.Header)
		})
	}
}
