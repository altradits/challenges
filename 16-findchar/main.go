package main

import "fmt"

// FindChar returns the index of the first occurrence of c in s.
// If c is not found, return -1.
func FindChar(s string, c rune) int {
	// TODO: Implement this function
	// Hint: Use for...range to get both index and character.
	// When you find a match, return the index immediately.
	// If the loop finishes, return -1.
	return -1
}

func main() {
	fmt.Println(FindChar("Hello", 'l'))
	fmt.Println(FindChar("Hello", 'H'))
	fmt.Println(FindChar("Hello", 'z'))
	fmt.Println(FindChar("", 'a'))
}
