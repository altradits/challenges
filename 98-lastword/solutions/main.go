package main

import (
	"fmt"
)

func LastWord(s string) string {
	// Trim trailing spaces
	end := len(s)
	for end > 0 && s[end-1] == ' ' {
		end--
	}

	// Find the start of the last word
	start := end
	for start > 0 && s[start-1] != ' ' {
		start--
	}

	return s[start:end] + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}