package piscine

func ReplaceChar(s string, old, new rune) string {
	result := ""
	for _, c := range s {
		if c == old {
			result += string(new)
		} else {
			result += string(c)
		}
	}
	return result
}
