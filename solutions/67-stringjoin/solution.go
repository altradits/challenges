package piscine

import "strings"

func Join(words []string, sep string) string {
	var b strings.Builder
	for i, w := range words {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(w)
	}
	return b.String()
}
