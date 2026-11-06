package piscine

import (
	"strings"
)

func LongestWord(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for _, word := range words[1:] {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}
