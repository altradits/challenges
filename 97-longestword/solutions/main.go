package main

import (
	"fmt"
	"strings"
)

func LongestWord(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for _, word := range words[1:] {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func main() {
	fmt.Println(LongestWord("Hello World"))
	fmt.Println(LongestWord("Go is fun"))
	fmt.Println(LongestWord(""))
	fmt.Println(LongestWord("a bb ccc dddd"))
}
