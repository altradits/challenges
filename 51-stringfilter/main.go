package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsVowel(r rune) bool {
	return strings.ContainsRune("aeiouAEIOU", r)
}

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsUpper(r rune) bool {
	return unicode.IsUpper(r)
}

func Filter(s string, f func(rune) bool) string {
	var b strings.Builder
	for _, r := range s {
		if f(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func main() {
	fmt.Println(Filter("hello", IsVowel))
	fmt.Println(Filter("abc123", IsDigit))
	fmt.Println(Filter("Go Lang", IsUpper))
	fmt.Println(Filter("", IsVowel))
}
