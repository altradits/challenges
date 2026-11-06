package piscine

func LastChar(s string) string {
	if s == "" {
		return ""
	}
	for _, c := range s {
		// Keep overwriting until the last character
		_ = c
	}
	// Get the last rune
	runes := []rune(s)
	return string(runes[len(runes)-1])
}
