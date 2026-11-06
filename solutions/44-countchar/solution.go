package piscine

func CountChar(s string, c rune) int {
	count := 0
	for _, ch := range s {
		if ch == c {
			count++
		}
	}
	return count
}
