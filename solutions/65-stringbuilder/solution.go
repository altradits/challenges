package piscine

import "strings"

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func RemoveVowels(s string) string {
	var b strings.Builder
	for _, c := range s {
		if !isVowel(c) {
			b.WriteRune(c)
		}
	}
	return b.String()
}
