package piscine

import "strings"

func Union(s1, s2 string) string {
	seen := make(map[rune]bool)
	var b strings.Builder
	for _, c := range s1 + s2 {
		if !seen[c] {
			seen[c] = true
			b.WriteRune(c)
		}
	}
	return b.String()
}
