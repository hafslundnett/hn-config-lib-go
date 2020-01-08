package libhttp

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_NewClient(t *testing.T) {
	// Test with pool error
	_, err := NewClient(mock.File)
	assert.Err(t, err, "while loading CA cert from file: failed to read CA file")

	// Test with no pool
	_, err = NewClient()
	assert.NoErr(t, err)
}

func Test_Do(t *testing.T) {
	client, err := NewClient()
	assert.NoErr(t, err)

	req, err := http.NewRequest(http.MethodGet, mock.URL, nil)
	assert.NoErr(t, err)

	err = client.Do(req)
	assert.Err(t, err, "while do-ing http request")
}

func Test_PostForm(t *testing.T) {
	client, err := NewClient()
	assert.NoErr(t, err)

	form := url.Values{}
	form.Add("addr", mock.Addr)
	form.Add("path", mock.Path)

	err = client.PostForm(mock.URL, form)
	assert.Err(t, err, "while post-ing http request")
}

func Test_checkRespCode(t *testing.T) {
	tests := []struct {
		name      string
		r         *http.Response
		wantErr   bool
		errWanted string
	}{
		{
			name:      "response with error",
			r:         &http.Response{StatusCode: 404},
			wantErr:   true,
			errWanted: "http error, status 404",
		}, {
			name:    "response without error",
			r:       &http.Response{StatusCode: 200},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkRespCode(tt.r)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}
}
