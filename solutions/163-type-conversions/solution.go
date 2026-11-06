package piscine

import "unicode/utf8"

func IntToFloat(n int) float64 {
	return float64(n)
}

func FloatToInt(f float64) int {
	return int(f)
}

func RuneToString(r rune) string {
	return string(r)
}

func StringToRunes(s string) []rune {
	return []rune(s)
}

func BytesToString(b []byte) string {
	return string(b)
}

func CountRunes(s string) int {
	return utf8.RuneCountInString(s)
}
