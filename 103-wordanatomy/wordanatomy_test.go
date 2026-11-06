package piscine

import "testing"

func TestWordAnatomy(t *testing.T) {
	cases := []struct {
		name      string
		word      string
		pre, root string
		suf       string
	}{
		{"prefix only", "unhappy", "un", "happy", ""},
		{"suffix only", "runner", "", "runn", "er"},
		{"both", "redoing", "re", "do", "ing"},
		{"neither", "planet", "", "planet", ""},
		{"mis", "misunderstand", "mis", "understand", ""},
		{"ful", "careful", "", "care", "ful"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pre, root, suf := WordAnatomy(tc.word)
			if pre != tc.pre || root != tc.root || suf != tc.suf {
				t.Errorf("WordAnatomy(%q) = (%q, %q, %q), want (%q, %q, %q)",
					tc.word, pre, root, suf, tc.pre, tc.root, tc.suf)
			}
		})
	}
}
