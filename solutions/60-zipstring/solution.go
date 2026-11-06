package piscine

import (
	"strconv"
	"strings"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	runes := []rune(s)
	count := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			b.WriteString(strconv.Itoa(count))
			b.WriteRune(runes[i-1])
			count = 1
		}
	}
	b.WriteString(strconv.Itoa(count))
	b.WriteRune(runes[len(runes)-1])
	return b.String()
}
