package piscine

import "strings"

func Split(s, sep string) []string {
	if sep == "" {
		out := make([]string, 0, len(s))
		for _, c := range s {
			out = append(out, string(c))
		}
		return out
	}
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
