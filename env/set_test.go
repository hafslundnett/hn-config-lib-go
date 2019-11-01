package env

import (
	"os"
	"testing"

	"github.com/giert/silver-octo-packet/assert"
	"github.com/giert/silver-octo-packet/mock"
)

func Test_Set(t *testing.T) {
	fields := []string{
		"MOCK_ADDR",
		"MOCK_TOKEN",
		"MOCK_ROLE",
		"MOCK_FILE",
		"MOCK_PATH",
	}

	tests := []struct {
		vars map[string]string
	}{
		{
			vars: map[string]string{
				fields[0]: mock.Addr,
				fields[1]: mock.Token,
				fields[2]: mock.Role,
				fields[3]: mock.File,
				fields[4]: mock.Path,
			},
		}, {
			vars: map[string]string{
				fields[0]: "",
				fields[1]: "",
				fields[2]: "",
				fields[3]: "",
				fields[4]: "",
			},
		},
	}

	for _, test := range tests {
		SetMap(test.vars)
		for _, field := range fields {
			assert.Result(t, os.Getenv(field), test.vars[field])
		}
	}

	for _, test := range tests {
		s := mapToSlice(test.vars)
		Set(s...)
		for _, field := range fields {
			assert.Result(t, os.Getenv(field), test.vars[field])
		}
	}
}

func mapToSlice(m map[string]string) []string {
	var s []string
	for k, v := range m {
		s = append(s, k, v)
	}
	return s
}
