package main

import "fmt"

// RetainFirstHalf returns the first half of a string.
// If the length is odd, round down. If empty or 1 char, return as-is.
func RetainFirstHalf(s string) string {
	// TODO: Implement this function
	// Hint: Check if len(s) <= 1 first.
	// Otherwise, return s[:len(s)/2]
	return ""
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
