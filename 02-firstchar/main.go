package main

import "fmt"

// FirstChar returns the first character of a string as a string.
// If the string is empty, return an empty string.
func FirstChar(s string) string {
	// TODO: Implement this function
	// Hint: Check if the string is empty first.
	// If not empty, use for...range to get the first rune,
	// then convert it to a string.
	return ""
}

func main() {
	fmt.Println(FirstChar("Hello"))
	fmt.Println(FirstChar(""))
	fmt.Println(FirstChar("G"))
	fmt.Println(FirstChar("Go is fun"))
}
