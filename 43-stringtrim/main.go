package main

import (
	"fmt"
	"unicode"
)

func Trim(s string) string {
	start := 0
	end := len(s)
	for start < end && unicode.IsSpace(rune(s[start])) {
		start++
	}
	for end > start && unicode.IsSpace(rune(s[end-1])) {
		end--
	}
	return s[start:end]
}

func main() {
	fmt.Printf("%q\n", Trim("  hello  "))
	fmt.Printf("%q\n", Trim("\t\nGo\n\t"))
	fmt.Printf("%q\n", Trim("no spaces"))
	fmt.Printf("%q\n", Trim("   "))
}
