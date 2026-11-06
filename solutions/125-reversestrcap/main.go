package main

import (
	"fmt"
	"os"
	"unicode"
)

// reverseCap lowercases every letter and uppercases the last letter of each
// word (the letter before a space or end of string). Spacing is preserved.
func reverseCap(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if !unicode.IsLetter(c) {
			continue
		}
		if i+1 == len(runes) || unicode.IsSpace(runes[i+1]) {
			runes[i] = unicode.ToUpper(c)
		} else {
			runes[i] = unicode.ToLower(c)
		}
	}
	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(reverseCap(os.Args[i]))
	}
}
