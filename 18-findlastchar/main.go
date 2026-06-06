package main

import "fmt"

// FindLastChar returns the index of the last occurrence of c in s.
// If c is not found, return -1.
func FindLastChar(s string, c rune) int {
	// TODO: Implement this function
	// Hint: Keep track of the last index where you found a match.
	// Initialize to -1, update whenever you find a match.
	// After the loop, return the last index (or -1 if never found).
	return -1
}

func main() {
	fmt.Println(FindLastChar("Hello", 'l'))
	fmt.Println(FindLastChar("Hello", 'H'))
	fmt.Println(FindLastChar("Hello", 'z'))
	fmt.Println(FindLastChar("banana", 'a'))
}
