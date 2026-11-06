package piscine

import "strings"

func ThirdChar(s string) string {
	var b strings.Builder
	i := 0
	for _, c := range s {
		i++
		if i%3 == 0 {
			b.WriteRune(c)
		}
	}
	b.WriteByte('\n')
	return b.String()
}
