package piscine

import "strings"

// Filter keeps the characters for which f returns true.
func Filter(s string, f func(rune) bool) string {
	var b strings.Builder
	for _, c := range s {
		if f(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}
