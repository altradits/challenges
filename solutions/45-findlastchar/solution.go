package piscine

func FindLastChar(s string, c rune) int {
	lastIndex := -1
	for i, ch := range s {
		if ch == c {
			lastIndex = i
		}
	}
	return lastIndex
}
