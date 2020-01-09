package hid

import (
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/env"
	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
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

func Test_New(t *testing.T) {
	err := env.Save(testenv...)
	assert.NoErr(t, err)

	tests := []struct {
		name      string
		want      string
		envslice  []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "no environment variables",
			wantErr:   true, // Fails successfully if HID is not running locally
			errWanted: "missing env var HID_ADDR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnv(t, tt.envslice)

			_, err := New()
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			err = env.Reset()
			assert.NoErr(t, err)
		})
	}

	err = env.Reset()
	assert.NoErr(t, err)
}
