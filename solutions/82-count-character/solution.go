package piscine

func CountChar(str string, c rune) int {
	if len(str) == 0 {
		return 0
	}

	count := 0
	for _, char := range str {
		if char == c {
			count++
		}
	}
	return count
}
