package piscine

import (
	"reflect"
	"testing"
)

func TestSplitJoin(t *testing.T) {
	splitCases := []struct {
		name string
		s    string
		sep  string
		want []string
	}{
		{"commas", "a,b,c", ",", []string{"a", "b", "c"}},
		{"space", "Hello World", " ", []string{"Hello", "World"}},
		{"empty sep", "abc", "", []string{"a", "b", "c"}},
		{"empty string", "", ",", []string{""}},
	}
	for _, tc := range splitCases {
		t.Run("split/"+tc.name, func(t *testing.T) {
			if got := Split(tc.s, tc.sep); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Split(%q, %q) = %q, want %q", tc.s, tc.sep, got, tc.want)
			}
		})
	}
	joinCases := []struct {
		name  string
		elems []string
		sep   string
		want  string
	}{
		{"commas", []string{"a", "b", "c"}, ",", "a,b,c"},
		{"space", []string{"Hello", "World"}, " ", "Hello World"},
		{"empty sep", []string{"a", "b", "c"}, "", "abc"},
		{"single empty", []string{""}, ",", ""},
		{"empty slice", []string{}, ",", ""},
	}
	for _, tc := range joinCases {
		t.Run("join/"+tc.name, func(t *testing.T) {
			if got := Join(tc.elems, tc.sep); got != tc.want {
				t.Errorf("Join(%q, %q) = %q, want %q", tc.elems, tc.sep, got, tc.want)
			}
		})
	}
}
