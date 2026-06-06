package main

import "fmt"

// ReplaceChar returns a new string where all occurrences of old are replaced with new.
func ReplaceChar(s string, old, new rune) string {
	// TODO: Implement this function
	// Hint: Iterate through the string.
	// When you find old, append new to the result.
	// Otherwise, append the current character.
	return ""
}

func main() {
	fmt.Println(ReplaceChar("Hello", 'l', 'L'))
	fmt.Println(ReplaceChar("Hello", 'z', 'X'))
	fmt.Println(ReplaceChar("", 'a', 'b'))
	fmt.Println(ReplaceChar("aaaa", 'a', 'b'))
}
