package piscine

import "testing"

func TestThirdTimeIsACharm(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"digits", "123456789", "369\n"},
		{"empty", "", "\n"},
		{"spaces", "a b c d e f", "b e\n"},
		{"short", "12", "\n"},
		{"exact", "abc", "c\n"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ThirdTimeIsACharm(tc.input); got != tc.want {
				t.Errorf("ThirdTimeIsACharm(%q) = %q, want %q (every 3rd char + newline)", tc.input, got, tc.want)
			}
		})
	}
}
