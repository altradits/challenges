package piscine

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		name string
		s    string
		sep  string
		want []string
	}{
		{"commas", "a,b,c", ",", []string{"a", "b", "c"}},
		{"space", "Hello World", " ", []string{"Hello", "World"}},
		{"empty sep", "abc", "", []string{"a", "b", "c"}},
		{"empty string", "", ",", []string{""}},
		{"no sep", "abc", "x", []string{"abc"}},
		{"consecutive", "a,,b", ",", []string{"a", "", "b"}},
		{"edge seps", ",a,", ",", []string{"", "a", ""}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Split(%q, %q) = %q, want %q", tc.s, tc.sep, got, tc.want)
			}
		})
	}
}
