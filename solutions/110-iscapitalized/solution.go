package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	inWord := false
	for _, c := range s {
		if c == ' ' {
			inWord = false
		} else if !inWord {
			inWord = true
			if c >= 'a' && c <= 'z' {
				return false
			}
		}
	}
	return true
}
