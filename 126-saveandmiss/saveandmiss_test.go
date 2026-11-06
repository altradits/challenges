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
		{"zero keeps", "hello you all ! ", 0, "hello you all ! "},
		{"zero keeps 2", "what is your name?", 0, "what is your name?"},
		{"negative keeps", "go Exercise Save and Miss", -5, "go Exercise Save and Miss"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SaveAndMiss(tc.s, tc.n); got != tc.want {
				t.Errorf("SaveAndMiss(%q, %d) = %q, want %q", tc.s, tc.n, got, tc.want)
			}
		})
	}
}
