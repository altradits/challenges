package piscine

func LongestWord(s string) string {
	words := fields(s)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for _, word := range words[1:] {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func fields(s string) []string {
	var words []string
	var current []byte
	inWord := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if inWord {
				words = append(words, string(current))
				current = nil
				inWord = false
			}
		} else {
			current = append(current, c)
			inWord = true
		}
	}

	if inWord {
		words = append(words, string(current))
	}

	return words
}
