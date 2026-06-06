package main

import "fmt"

func WordCount(s string) int {
	// TODO: Implement this function
	// Hint: Iterate through the string character by character.
	// Track whether you are currently inside a word.
	// When you transition from a space to a non-space, increment the count.
	return 0
}

func main() {
	fmt.Println(WordCount("Hello World"))
	fmt.Println(WordCount("Go is fun!"))
	fmt.Println(WordCount(""))
	fmt.Println(WordCount("   "))
	fmt.Println(WordCount("One  two   three"))
}
