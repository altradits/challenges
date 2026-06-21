package main

import (
	"fmt"
	"os"
	"strings"
)

func toPigLatin(word string) string {
	vowels := "aeiouAEIOU"
	for i, c := range word {
		if strings.ContainsRune(vowels, c) {
			if i == 0 {
				return word + "ay"
			}
			return word[i:] + word[:i] + "ay"
		}
	}
	return "No vowels"
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := strings.Fields(os.Args[1])
	if len(words) != 1 {
		return
	}
	fmt.Println(toPigLatin(words[0]))
}
