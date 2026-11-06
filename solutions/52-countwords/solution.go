package piscine

func isWordChar(c rune) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9'
}

func CountWordsAdvanced(s string) int {
	count := 0
	inWord := false
	for _, c := range s {
		if isWordChar(c) {
			if !inWord {
				count++
				inWord = true
			}
		} else {
			inWord = false
		}
	}
	return count
}
