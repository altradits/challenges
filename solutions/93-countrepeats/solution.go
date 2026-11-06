package piscine

func CountRepeats(s string) int {
	if len(s) < 2 {
		return 0
	}

	count := 0
	inRepeat := false

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if !inRepeat {
				count++
				inRepeat = true
			}
		} else {
			inRepeat = false
		}
	}

	return count
}
