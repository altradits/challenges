package piscine

import "strings"

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// CamelToSnakeCase inserts '_' before each uppercase letter. Inputs that are
// not plain camelCase (digits, symbols, or consecutive capitals) are invalid
// and returned unchanged.
func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	prevUpper := false
	for _, r := range runes {
		if !isUpper(r) && !isLower(r) {
			return s
		}
		if isUpper(r) {
			if prevUpper {
				return s
			}
			prevUpper = true
		} else {
			prevUpper = false
		}
	}
	var b strings.Builder
	for i, r := range runes {
		if isUpper(r) && i > 0 {
			b.WriteByte('_')
		}
		b.WriteRune(r)
	}
	return b.String()
}
