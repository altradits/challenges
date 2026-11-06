package piscine

import "strings"

func Map(s string, f func(rune) rune) string {
	var b strings.Builder
	for _, c := range s {
		b.WriteRune(f(c))
	}
	return b.String()
}
