package piscine

import "strings"

func WordFrequency(s string) map[string]int {
	freq := make(map[string]int)
	for _, w := range strings.Fields(s) {
		freq[w]++
	}
	return freq
}
