package piscine

import "testing"

func TestThirdChar(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"digits", "123456789", "369\n"},
		{"empty", "", "\n"},
		{"with spaces", "a b c d e f", "b e\n"},
		{"short", "12", "\n"},
		{"exact three", "abc", "c\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ThirdChar(tc.input); got != tc.want {
				t.Errorf("ThirdChar(%q) = %q, want %q (every 3rd char + newline)", tc.input, got, tc.want)
			}
		})
	}
}
