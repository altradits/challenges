package piscine

import "strings"

func WordCount(s string) int {
	return len(strings.Fields(s))
}
