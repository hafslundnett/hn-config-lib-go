package util

import (
	"reflect"
	"testing"
)

func TestMapToSlice(t *testing.T) {
	tests := []struct {
		name  string
		arg   map[interface{}]interface{}
		want1 []interface{}
		want2 []interface{}
	}{
		{
			name:  "only strings",
			arg:   map[interface{}]interface{}{"a": "1", "b": "2"},
			want1: []interface{}{"a", "1", "b", "2"},
			want2: []interface{}{"b", "2", "a", "1"},
		},
		{
			name:  "mixed types",
			arg:   map[interface{}]interface{}{1: true, 2: false},
			want1: []interface{}{1, true, 2, false},
			want2: []interface{}{2, false, 1, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToSlice(tt.arg); !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("MapToSlice() = %v, want %v or %v", got, tt.want1, tt.want2)
			}
		})
	}
}
