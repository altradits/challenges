package piscine

import "testing"

func TestRemoveVowels(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{"hello", "Hello World", "Hll Wrld"},
		{"all vowels", "AEIOUaeiou", ""},
		{"go", "Go Programming", "G Prgrmmng"},
		{"empty", "", ""},
		{"no vowels", "rhythm", "rhythm"},
		{"y stays", "sky", "sky"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RemoveVowels(tc.input); got != tc.want {
				t.Errorf("RemoveVowels(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
