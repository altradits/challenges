package piscine

func TitleCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	capitalizeNext := true

	for i, c := range runes {
		if c == ' ' {
			capitalizeNext = true
		} else if capitalizeNext {
			if c >= 'a' && c <= 'z' {
				runes[i] = c - 'a' + 'A'
			}
			capitalizeNext = false
		} else {
			if c >= 'A' && c <= 'Z' {
				runes[i] = c - 'A' + 'a'
			}
		}
	}

	return string(runes)
}
