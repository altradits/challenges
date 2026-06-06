package main

import "fmt"

// IsEmpty returns true if the string is empty, false otherwise.
// Do NOT use len(). Instead, iterate and check if any character exists.
func IsEmpty(s string) bool {
	// TODO: Implement this function
	// Hint: Use for...range to iterate.
	// If you find at least one character, return false immediately.
	// If the loop finishes without finding any, return true.
	return false
}

func main() {
	fmt.Println(IsEmpty("Hello"))
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty(" "))
	fmt.Println(IsEmpty("G"))
}
