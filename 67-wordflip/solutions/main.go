package main

import (
	"fmt"
	"strings"
)

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}

	// Trim spaces and split into words
	words := strings.Fields(str)
	if len(words) == 0 {
		return "\n"
	}

	// Reverse the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
