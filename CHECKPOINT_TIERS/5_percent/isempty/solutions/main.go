package main

import "fmt"

func IsEmpty(s string) bool {
	// TODO: Implement this function
	// Hint: Iterate through the string using for...range.
	// If you find at least one character, return false.
	// If the loop finishes without finding any, return true.
	return false
}

func main() {
	fmt.Println(IsEmpty("Hello"))
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty(" "))
	fmt.Println(IsEmpty("G"))
}
