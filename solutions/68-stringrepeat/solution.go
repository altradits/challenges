package piscine

import "strings"

func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}
