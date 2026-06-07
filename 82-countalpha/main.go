package main

import "fmt"

// CountAlpha returns the number of alphabetic characters (a-z, A-Z) in a string.
func CountAlpha(s string) int {
	// TODO: Implement this function
	// Hint: Check if each character is between 'A'-'Z' OR 'a'-'z'.
	// Use the OR operator (||) to combine both conditions.
	return 0
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
