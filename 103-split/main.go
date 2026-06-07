package main

import "fmt"

// Split splits a string by a separator and returns a slice of substrings.
// If sep is empty, split by individual characters.
func Split(s, sep string) []string {
	// TODO: Implement this function
	// Hint: Track the start position of each segment.
	// When you find the separator, extract the substring and add to result.
	return nil
}

func main() {
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("Hello World", " "))
	fmt.Printf("%q\n", Split("abc", ""))
	fmt.Printf("%q\n", Split("", ","))
}
