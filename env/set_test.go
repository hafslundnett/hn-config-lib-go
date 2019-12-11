package env

import (
	"os"
	"testing"

	"github.com/hafslundnett/hn-config-lib-go/testing/assert"
	"github.com/hafslundnett/hn-config-lib-go/testing/mock"
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
			nil,
		},
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
		err := SetMap(test.vars)
		assert.NoErr(t, err)

		for _, field := range fields {
			assert.Result(t, os.Getenv(field), test.vars[field])
		}
	}

	for _, test := range tests {
		s := mapToSlice(test.vars)
		err := Set(s...)
		assert.NoErr(t, err)

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
