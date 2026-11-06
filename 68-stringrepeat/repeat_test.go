package piscine

import "testing"

func TestRepeat(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{"three", "abc", 3, "abcabcabc"},
		{"five", "x", 5, "xxxxx"},
		{"zero", "go", 0, ""},
		{"empty", "", 10, ""},
		{"negative", "go", -2, ""},
		{"one", "go", 1, "go"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Repeat(tc.s, tc.n); got != tc.want {
				t.Errorf("Repeat(%q, %d) = %q, want %q", tc.s, tc.n, got, tc.want)
			}
		})
	}
}
