package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Map(s string, f func(rune) rune) string {
	var b strings.Builder
	for _, r := range s {
		b.WriteRune(f(r))
	}
	return b.String()
}

func main() {
	fmt.Println(Map("abc", unicode.ToUpper))
	fmt.Println(Map("Hello", unicode.ToLower))
	fmt.Println(Map("123", func(r rune) rune { return r + 1 }))
	fmt.Println(Map("", func(r rune) rune { return r }))
}
