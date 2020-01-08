package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

var testenv = []string{
	envars["addr"],
	envars["github"],
	envars["account"],
	envars["mount"],
	envars["role"],
}

func Test_Vault_Configure(t *testing.T) {
	err := env.Save(testenv...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		want      Config
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			want:      Config{},
			wantErr:   true,
			errWanted: "missing env var " + envars["addr"],
		}, {
			name:     "github environment variables",
			envslice: []string{envars["addr"], mock.Addr, envars["github"], mock.Token},
			want:     Config{Addr: mock.Addr, GithubToken: mock.Token, Client: mock.Client},
			wantErr:  false,
		}, {
			name:     "k8 environment variables",
			envslice: []string{envars["addr"], mock.Addr, envars["mount"], mock.Path, envars["account"], mock.Path, envars["role"], mock.Role},
			want:     Config{Addr: mock.Addr, K8ServicePath: mock.Path, K8MountPath: mock.Path, K8Role: mock.Role, Client: mock.Client},
			wantErr:  false,
		}, {
			name:      "k8 no role variable",
			envslice:  []string{envars["addr"], mock.Addr, envars["mount"], mock.Path, envars["account"], mock.Path},
			want:      Config{},
			wantErr:   true,
			errWanted: "missing either ROLE or GITHUB_TOKEN env var",
		}, {
			name:     "k8 default values",
			envslice: []string{envars["addr"], mock.Addr, envars["role"], mock.Role},
			want:     Config{Addr: mock.Addr, K8ServicePath: "/var/run/secrets/kubernetes.io/serviceaccount/token", K8MountPath: "kubernetes", K8Role: mock.Role, Client: mock.Client},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnv(t, tt.envslice)

			vault := new(Vault)
			err = vault.Configure(mock.Client)

			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			assert.Result(t, vault.Config, tt.want)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
