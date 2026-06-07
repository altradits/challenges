package main

import (
	"fmt"
)

func FirstWord(s string) string {
	// Find the start of the first word (skip leading spaces)
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}

	// Find the end of the first word (stop at space or end)
	end := start
	for end < len(s) && s[end] != ' ' {
		end++
	}

	return s[start:end] + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}