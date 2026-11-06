package piscine

import "testing"

func TestJoinWords(t *testing.T) {
	cases := []struct {
		name  string
		words []string
		sep   string
		want  string
	}{
		{"dashes", []string{"a", "b", "c"}, "-", "a-b-c"},
		{"space", []string{"hello", "world"}, " ", "hello world"},
		{"empty", []string{}, ",", ""},
		{"single", []string{"single"}, "|", "single"},
		{"nil", nil, ",", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Join(tc.words, tc.sep); got != tc.want {
				t.Errorf("Join(%q, %q) = %q, want %q", tc.words, tc.sep, got, tc.want)
			}
		})
	}
}
