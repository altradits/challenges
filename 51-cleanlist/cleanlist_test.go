package piscine

import (
	"reflect"
	"testing"
)

func TestCleanList(t *testing.T) {
	cases := []struct {
		name  string
		input []string
		want  []string
	}{
		{"trim", []string{"Hello", " World ", "Go"}, []string{"Hello", "World", "Go"}},
		{"drop empty", []string{"", "a", "", "b"}, []string{"a", "b"}},
		{"empty slice", []string{}, nil},
		{"only spaces", []string{"  "}, nil},
		{"nil", nil, nil},
		{"tabs", []string{"\tx\t"}, []string{"x"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CleanList(tc.input)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CleanList(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
