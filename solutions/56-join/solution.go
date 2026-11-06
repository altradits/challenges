package piscine

import "strings"

func Join(elems []string, sep string) string {
	var b strings.Builder
	for i, e := range elems {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(e)
	}
	return b.String()
}
