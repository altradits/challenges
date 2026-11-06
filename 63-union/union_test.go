package piscine

import "testing"

func TestUnion(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 string
		want   string
	}{
		{"piscine", "zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj", "zpadintoqefwjy"},
		{"digits", "ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd", "df6vewg4thras"},
		{"empty", "", "", ""},
		{"one empty", "abc", "", "abc"},
		{"dupes collapse", "aaa", "aa", "a"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Union(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Union(%q, %q) = %q, want %q (unique chars, first-appearance order)", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
