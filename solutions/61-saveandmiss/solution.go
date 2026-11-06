package piscine

import "strings"

func SaveAndMiss(s string, n int) string {
	if n <= 0 {
		return s
	}
	var b strings.Builder
	runes := []rune(s)
	for i := 0; i < len(runes); {
		for k := 0; k < n && i < len(runes); k++ {
			b.WriteRune(runes[i])
			i++
		}
		i += n
	}
	return b.String()
}
