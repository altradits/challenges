package main

import (
	"fmt"
	"strings"
)

func RemoveVowels(s string) string {
	var b strings.Builder
	vowels := "aeiouAEIOU"
	for _, r := range s {
		if !strings.ContainsRune(vowels, r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func main() {
	fmt.Println(RemoveVowels("Hello World"))
	fmt.Println(RemoveVowels("AEIOUaeiou"))
	fmt.Println(RemoveVowels("Go Programming"))
	fmt.Println(RemoveVowels(""))
}
