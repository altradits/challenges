package main

import "fmt"

func FindSubstring(text, substring string) int {
	// TODO: Implement this function
	// Hint: Use a nested loop approach.
	// 1. Handle edge cases (empty substring, substring longer than text).
	// 2. Loop through each possible starting position in text.
	// 3. For each position, check if substring matches.
	return -1
}

func main() {
	fmt.Println(FindSubstring("Hello World", "World"))
	fmt.Println(FindSubstring("Hello World", "hello"))
	fmt.Println(FindSubstring("Hello World", ""))
	fmt.Println(FindSubstring("Hello World", "xyz"))
	fmt.Println(FindSubstring("banana", "ana"))
}
