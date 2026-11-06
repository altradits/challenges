package piscine

import (
	"strings"
	"unicode"
)

func ReverseStrCap(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		runes := []rune(strings.ToLower(w))
		if len(runes) > 0 {
			runes[len(runes)-1] = unicode.ToUpper(runes[len(runes)-1])
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}
