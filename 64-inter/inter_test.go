package piscine

import "testing"

func TestIntersection(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 string
		want   string
	}{
		{"piscine", "padinton", "paqefwtdjetyiytjneytjoeyjnejeyj", "padinto"},
		{"digits", "ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd", "df6ewg4"},
		{"empty", "", "", ""},
		{"disjoint", "abc", "xyz", ""},
		{"dupes collapse", "aabb", "ab", "ab"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Intersection(tc.s1, tc.s2); got != tc.want {
				t.Errorf("Intersection(%q, %q) = %q, want %q (shared chars, s1 order)", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
