package main

import "fmt"

// LongestWord returns the longest word in a string.
// If there are multiple words with the same maximum length, return the first one.
// A word is a sequence of non-space characters.
func LongestWord(s string) string {
	// TODO: Implement this function
	// Hint: Split the string into words, then find the longest one.
	// Track the longest word and its length as you iterate.
	return ""
}

func main() {
	fmt.Println(LongestWord("Hello World"))
	fmt.Println(LongestWord("Go is fun"))
	fmt.Println(LongestWord(""))
	fmt.Println(LongestWord("a bb ccc dddd"))
}
