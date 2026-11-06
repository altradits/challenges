package piscine

import "testing"

func TestJoin(t *testing.T) {
	cases := []struct {
		name  string
		elems []string
		sep   string
		want  string
	}{
		{"commas", []string{"a", "b", "c"}, ",", "a,b,c"},
		{"space", []string{"Hello", "World"}, " ", "Hello World"},
		{"empty sep", []string{"a", "b", "c"}, "", "abc"},
		{"single empty", []string{""}, ",", ""},
		{"empty slice", []string{}, ",", ""},
		{"nil slice", nil, ",", ""},
		{"multi sep", []string{"a", "b"}, "<->", "a<->b"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Join(tc.elems, tc.sep); got != tc.want {
				t.Errorf("Join(%q, %q) = %q, want %q", tc.elems, tc.sep, got, tc.want)
			}
		})
	}
}
