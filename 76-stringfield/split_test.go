package piscine

import (
	"reflect"
	"testing"
)

func TestSplitFields(t *testing.T) {
	cases := []struct {
		name string
		s    string
		sep  string
		want []string
	}{
		{"commas", "a,b,c", ",", []string{"a", "b", "c"}},
		{"colons", "a:b:c", ":", []string{"a", "b", "c"}},
		{"no sep", "abc", "x", []string{"abc"}},
		{"empty s", "", ",", []string{""}},
		{"consecutive", "a,,b", ",", []string{"a", "", "b"}},
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
