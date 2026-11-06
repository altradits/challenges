package piscine

func FindSubstring(text, substring string) int {
	if len(substring) == 0 {
		return 0
	}
	if len(substring) > len(text) {
		return -1
	}
	for i := 0; i <= len(text)-len(substring); i++ {
		match := true
		for j := 0; j < len(substring); j++ {
			if text[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
