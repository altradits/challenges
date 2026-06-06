package main

import "fmt"

// WordCount returns the number of words in a string.
// A word is a sequence of non-space characters separated by spaces.
func WordCount(s string) int {
	// TODO: Implement this function
	// Hint: Track whether you are currently inside a word.
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
