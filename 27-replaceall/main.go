package main

import "fmt"

// ReplaceAll returns a new string where all occurrences of old are replaced with new.
func ReplaceAll(text, old, new string) string {
	// TODO: Implement this function
	// Hint: Use a loop to scan through text.
	// When you find old at the current position, append new to the result.
	// Otherwise, append the current character.
	return ""
}

func main() {
	fmt.Println(ReplaceAll("Hello World", "l", "L"))
	fmt.Println(ReplaceAll("banana", "na", "NA"))
	fmt.Println(ReplaceAll("Hello World", "", "x"))
	fmt.Println(ReplaceAll("abcabc", "abc", "xyz"))
	fmt.Println(ReplaceAll("no match here", "z", "Z"))
}
