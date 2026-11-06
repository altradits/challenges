package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// cleanWord strips leading and trailing characters that are not letters or
// digits. Inner punctuation (like the apostrophe in "don't") is kept.
func cleanWord(w string) string {
	return strings.TrimFunc(w, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	pattern := os.Args[1]
	text := os.Args[2]
	if len(pattern) < 2 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
		return
	}
	inner := pattern[1 : len(pattern)-1]
	if inner == "" {
		return
	}
	alternatives := strings.Split(inner, "|")
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}
	count := 0
	for _, word := range words {
		for _, alt := range alternatives {
			if alt != "" && strings.Contains(word, alt) {
				count++
				fmt.Printf("%d: %s\n", count, cleanWord(word))
			}
		}
	}
}
