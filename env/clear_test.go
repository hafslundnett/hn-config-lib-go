package env

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
)

func Test_Clear(t *testing.T) {
	tests := []struct {
		name      string
		vars      []string
		vals      []string
		wantErr   bool
		errWanted string
	}{
		{
			name:      "blank strings",
			vars:      []string{""},
			vals:      []string{""},
			wantErr:   true,
			errWanted: "setenv",
		},
		{
			name:    "ordinary strings",
			vars:    []string{"CLEAR_TEST1", "CLEAR_TEST2"},
			vals:    []string{"value1", "value2"},
			wantErr: false,
		},
		{
			name:    "the same strings as lowercase",
			vars:    []string{"clear_test1", "clear_test2"},
			vals:    []string{"value1", "value2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.vars {
				err := os.Setenv(v, tt.vals[i])
				assert.WantErr(t, tt.wantErr, err, tt.errWanted)
			}

			Save(tt.vars...)

			err := Clear(tt.vars...)
			assert.WantErr(t, tt.wantErr, err, tt.errWanted)

			for _, v := range tt.vars {
				assert.Result(t, os.Getenv(v), "")
			}

			err = Reset()
			assert.NoErr(t, err)

			for i, v := range tt.vars {
				assert.Result(t, os.Getenv(v), tt.vals[i])
			}
		})
	}
}
