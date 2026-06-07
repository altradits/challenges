package main

import "fmt"

func SaveAndMiss(s string, n int) string {
	// TODO: Implement this function
	// Hint: Track position within 2*n cycle.
	// Save when pos < n, skip when pos >= n.
	return ""
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
