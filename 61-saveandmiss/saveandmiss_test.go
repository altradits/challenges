package piscine

import "testing"

func TestSaveAndMiss(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"digits", "123456789", 3, "123789"},
		{"alphabet", "abcdefghijklmnopqrstuvwyz", 3, "abcghimnostuz"},
		{"empty", "", 3, ""},
		{"zero keeps all", "hello", 0, "hello"},
		{"one", "abcdef", 1, "ace"},
		{"n bigger than string", "abc", 9, "abc"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SaveAndMiss(tc.s, tc.n); got != tc.want {
				t.Errorf("SaveAndMiss(%q, %d) = %q, want %q", tc.s, tc.n, got, tc.want)
			}
		})
	}
}
