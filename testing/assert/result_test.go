package assert

import (
	"testing"
)

func Test_Result(t *testing.T) {
	somestructs := []struct {
		id   int
		name string
	}{
		{24, "some"},
		{24, "some"},
	}
	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{
			name: "string",
			got:  "a",
			want: "a",
		}, {
			name: "bool",
			got:  true,
			want: true,
		}, {
			name: "int",
			got:  42,
			want: 42,
		}, {
			name: "struct",
			got:  somestructs[0],
			want: somestructs[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Result(t, tt.got, tt.want)
		})
	}
}
