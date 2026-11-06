package piscine

import (
	"strings"
	"unicode"
)

func SplitWords(s string) []string {
	out := []string{}
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 0 {
			out = append(out, cur.String())
			cur.Reset()
		}
	}
	for _, c := range s {
		if unicode.IsSpace(c) {
			flush()
		} else {
			cur.WriteRune(c)
		}
	}
	flush()
	return out
}
