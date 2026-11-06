package piscine

import "strings"

func ReplaceAll(text, old, new string) string {
	if old == "" {
		return text
	}
	var b strings.Builder
	for {
		i := strings.Index(text, old)
		if i < 0 {
			b.WriteString(text)
			return b.String()
		}
		b.WriteString(text[:i])
		b.WriteString(new)
		text = text[i+len(old):]
	}
}
