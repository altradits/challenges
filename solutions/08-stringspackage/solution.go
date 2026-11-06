package piscine

import "strings"

func SentenceStats(s string) (words int, chars int, hasHello bool) {
	fields := strings.Fields(s)
	words = len(fields)
	for _, r := range s {
		if !isSpace(r) {
			chars++
		}
	}
	for _, w := range fields {
		if strings.EqualFold(w, "hello") {
			hasHello = true
			break
		}
	}
	return words, chars, hasHello
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}
