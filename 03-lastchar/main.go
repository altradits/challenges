package main

import "fmt"

// LastChar returns the last character of a string as a string.
// If the string is empty, return an empty string.
func LastChar(s string) string {
	// TODO: Implement this function
	// Hint: Check if the string is empty first.
	// If not empty, calculate the last index (len(s) - 1),
	// get the character at that position, and convert to string.
	return ""
}

func main() {
	fmt.Println(LastChar("Hello"))
	fmt.Println(LastChar(""))
	fmt.Println(LastChar("G"))
	fmt.Println(LastChar("Go is fun"))
}
