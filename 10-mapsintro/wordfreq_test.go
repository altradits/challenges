package piscine

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{"repeat", "go is go", map[string]int{"go": 2, "is": 1}},
		{"single", "hello", map[string]int{"hello": 1}},
		{"empty", "", map[string]int{}},
		{"multi", "a b a c a", map[string]int{"a": 3, "b": 1, "c": 1}},
		{"extra spaces", "  go   go ", map[string]int{"go": 2}},
		{"case sensitive", "Go go GO", map[string]int{"Go": 1, "go": 1, "GO": 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WordFrequency(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("WordFrequency(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
