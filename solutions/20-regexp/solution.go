package piscine

import "regexp"

var (
	emailRe  = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	numberRe = regexp.MustCompile(`[0-9]+`)
)

func ValidateEmail(s string) bool {
	return emailRe.MatchString(s)
}

func ExtractNumbers(s string) []string {
	matches := numberRe.FindAllString(s, -1)
	if matches == nil {
		return []string{}
	}
	return matches
}

func MaskPhone(s string) string {
	if len(s) <= 4 {
		return s
	}
	masked := make([]byte, len(s))
	for i := 0; i < len(s)-4; i++ {
		masked[i] = '*'
	}
	copy(masked[len(s)-4:], s[len(s)-4:])
	return string(masked)
}
