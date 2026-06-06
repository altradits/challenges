package main

import "fmt"

// CleanList takes a slice of strings and returns a new slice with:
// - Leading and trailing spaces removed from each string
// - Empty strings removed
func CleanList(inputs []string) []string {
	// TODO: Implement this function
	// Hint: Create a new slice.
	// For each string, trim spaces and check if it's not empty.
	// Only add non-empty strings to the result.
	return nil
}

func main() {
	fmt.Println(CleanList([]string{"Hello", " World ", "Go"}))
	fmt.Println(CleanList([]string{"", "a", "", "b"}))
	fmt.Println(CleanList([]string{}))
	fmt.Println(CleanList([]string{"  "}))
}
