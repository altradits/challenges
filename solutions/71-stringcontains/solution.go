package piscine

import "strings"

func Contains(s, substr string) bool {
	return strings.Index(s, substr) >= 0
}
