package assert

import (
	"testing"
)

var somestructs = []struct {
	id   int
	name string
}{
	{24, "some"},
	{24, "some"},
}

func Test_Result(t *testing.T) {
	type args struct {
		t    *testing.T
		got  interface{}
		want interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"string", args{t: t, got: "a", want: "a"}},
		{"bool", args{t: t, got: true, want: true}},
		{"int", args{t: t, got: 42, want: 42}},
		{"struct", args{t: t, got: somestructs[0], want: somestructs[1]}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Result(tt.args.t, tt.args.got, tt.args.want)
		})
	}
}
