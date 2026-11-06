package piscine

import "strings"

func Intersection(s1, s2 string) string {
	inS2 := make(map[rune]bool)
	for _, c := range s2 {
		inS2[c] = true
	}
	seen := make(map[rune]bool)
	var b strings.Builder
	for _, c := range s1 {
		if inS2[c] && !seen[c] {
			seen[c] = true
			b.WriteRune(c)
		}
	}
	return b.String()
}
