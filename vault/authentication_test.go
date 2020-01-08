package vault

import (
	"io"
	"net/http"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

func Test_Vault_Authenticate(t *testing.T) {
	tests := []struct {
		name      string
		vault     Vault
		wantErr   bool
		errWanted string
	}{
		{
			name:    "Github authentication",
			vault:   Vault{Config: Config{GithubToken: mock.Token, Client: mock.Client}},
			wantErr: false,
		}, {
			name:      "K8 authentication",
			vault:     Vault{Config: Config{K8MountPath: mock.Path, K8ServicePath: mock.Path}},
			wantErr:   true,
			errWanted: "failed to read jwt token from " + mock.Path,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.vault.Authenticate()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}
}

func Test_authReq(t *testing.T) {
	tests := []struct {
		name      string
		cfg       Config
		want      *http.Request
		wantErr   bool
		errWanted string
	}{
		{
			name:    "Github request",
			cfg:     Config{Addr: mock.Addr, GithubToken: mock.Token},
			want:    mock.Request(t, "POST", mock.Addr+"/v1/auth/github/login", "{\"token\":\""+mock.Token+"\"}"),
			wantErr: false,
		}, {
			name:      "K8 request",
			cfg:       Config{Addr: mock.Addr, K8MountPath: mock.Path, K8ServicePath: mock.Path},
			wantErr:   true,
			errWanted: "failed to read jwt token from " + mock.Path,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authReq(tt.cfg)

			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			if !tt.wantErr {
				assert.Result(t, got.Method, tt.want.Method)
				assert.DeepResult(t, got.URL, tt.want.URL)
				assert.DeepResult(t, got.Body, tt.want.Body)
			}
		})
	}
}

func Test_authBody(t *testing.T) {
	tests := []struct {
		name      string
		cfg       Config
		want      io.Reader
		wantPath  string
		wantErr   bool
		errWanted string
	}{
		{
			name:     "Github login body",
			cfg:      Config{GithubToken: mock.Token},
			want:     mock.Body("{\"token\":\"" + mock.Token + "\"}"),
			wantPath: "auth/github/login",
			wantErr:  false,
		}, {
			name:      "K8 login body",
			cfg:       Config{K8MountPath: mock.Path, K8ServicePath: mock.Path},
			wantPath:  "auth/" + mock.Path + "/login",
			wantErr:   true,
			errWanted: "failed to read jwt token from " + mock.Path,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotPath, err := authBody(tt.cfg)

			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			assert.Result(t, gotPath, tt.wantPath)

			if !tt.wantErr {
				assert.DeepResult(t, got, tt.want)
			}
		})
	}
}

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
