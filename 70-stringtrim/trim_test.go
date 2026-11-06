package piscine

import "testing"

func TestTrim(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"spaces", "  hello  ", "hello"},
		{"tabs newlines", "\t\nGo\n\t", "Go"},
		{"inner kept", "no spaces", "no spaces"},
		{"all space", "   ", ""},
		{"empty", "", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Trim(tc.input); got != tc.want {
				t.Errorf("Trim(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
