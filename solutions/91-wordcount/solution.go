package piscine

func WordCount(s string) int {
	count := 0
	inWord := false

	for _, c := range s {
		if c == ' ' {
			inWord = false
		} else if !inWord {
			count++
			inWord = true
		}
	}

	return count
}
