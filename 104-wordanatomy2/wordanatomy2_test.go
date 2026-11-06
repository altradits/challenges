package piscine

import "testing"

func TestWordanatomy(t *testing.T) {
	pref := []string{"an", "en", "un"}
	suff := []string{"ible", "able"}
	cases := []struct {
		name string
		word string
		want string
	}{
		{"both", "understandable", `prefix: "un", suffix: "able"`},
		{"prefix only", "understand", `prefix: "un", suffix: ""`},
		{"neither", "apple", `prefix: "", suffix: ""`},
		{"suffix only", "eatable", `prefix: "", suffix: "able"`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Wordanatomy(tc.word, pref, suff); got != tc.want {
				t.Errorf("Wordanatomy(%q) = %s, want %s", tc.word, got, tc.want)
			}
		})
	}
}
