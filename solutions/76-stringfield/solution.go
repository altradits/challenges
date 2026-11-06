package piscine

import "strings"

func Split(s, sep string) []string {
	var out []string
	for {
		i := strings.Index(s, sep)
		if i < 0 {
			return append(out, s)
		}
		out = append(out, s[:i])
		s = s[i+len(sep):]
	}
}
