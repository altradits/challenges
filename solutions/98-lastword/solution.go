package piscine

func LastWord(s string) string {
	// Trim trailing spaces
	end := len(s)
	for end > 0 && s[end-1] == ' ' {
		end--
	}

	// Find the start of the last word
	start := end
	for start > 0 && s[start-1] != ' ' {
		start--
	}

	return s[start:end] + "\n"
}
