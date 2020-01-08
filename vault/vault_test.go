package vault

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
)

// replaceEnv compacts environment variables handling to increase readability of tests.
func replaceEnv(t *testing.T, vars []string) {
	t.Helper()

	err := env.Clear(testenv...)
	assert.NoErr(t, err)

	if vars != nil {
		err = env.Set(vars...)
		assert.NoErr(t, err)
	}
}

func TestNew(t *testing.T) {
	err := env.Save(testenv...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			envslice:  []string{},
			wantErr:   true,
			errWanted: "missing env var VAULT_ADDR",
		}, {
			name:      "broken authentification",
			envslice:  []string{envars["addr"], mock.Addr, envars["github"], mock.Token},
			wantErr:   true,
			errWanted: "while do-ing http request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnv(t, tt.envslice)

			_, err := New()

			assert.WantErr(t, tt.wantErr, err, tt.errWanted)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
