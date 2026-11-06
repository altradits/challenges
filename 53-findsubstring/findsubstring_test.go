package piscine

import "testing"

func TestFindSubstring(t *testing.T) {
	cases := []struct {
		name      string
		text, sub string
		want      int
	}{
		{"found", "Hello World", "World", 6},
		{"case sensitive", "Hello World", "hello", -1},
		{"empty sub", "Hello World", "", 0},
		{"missing", "Hello World", "xyz", -1},
		{"overlap start", "banana", "ana", 1},
		{"whole", "abc", "abc", 0},
		{"longer than text", "ab", "abc", -1},
		{"empty text", "", "a", -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindSubstring(tc.text, tc.sub); got != tc.want {
				t.Errorf("FindSubstring(%q, %q) = %d, want %d", tc.text, tc.sub, got, tc.want)
			}
		})
	}
}
