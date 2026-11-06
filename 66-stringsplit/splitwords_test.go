package piscine

import (
	"reflect"
	"testing"
)

func TestSplitWords(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  []string
	}{
		{"two", "Hello World", []string{"Hello", "World"}},
		{"messy", "  Go   is  fun  ", []string{"Go", "is", "fun"}},
		{"empty", "", []string{}},
		{"single", "single", []string{"single"}},
		{"tabs", "a\tb", []string{"a", "b"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitWords(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SplitWords(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
