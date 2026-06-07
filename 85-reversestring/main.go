package main

import "fmt"

// ReverseString returns the string reversed.
// Hint: Go strings are UTF-8 encoded. Use []rune to handle all characters.
func ReverseString(s string) string {
	// TODO: Implement this function
	// 1. Convert the string to a slice of runes.
	// 2. Swap elements from the start and end of the slice.
	// 3. Convert the rune slice back to a string.
	return ""
}

func main() {
	fmt.Println(ReverseString("Hello"))
	fmt.Println(ReverseString("Go is fun!"))
	fmt.Println(ReverseString(""))
	fmt.Println(ReverseString("a"))
}
