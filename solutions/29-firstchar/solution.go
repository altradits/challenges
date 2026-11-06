package piscine

func FirstChar(s string) string {
	if s == "" {
		return ""
	}
	for _, c := range s {
		return string(c)
	}
	return ""
}
