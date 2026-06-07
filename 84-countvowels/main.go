package main

import "fmt"

// CountVowels returns the number of vowels (a, e, i, o, u) in a string.
// Count both lowercase and uppercase vowels.
func CountVowels(s string) int {
	// TODO: Implement this function
	// Hint: Define what a vowel is (a, e, i, o, u - both cases).
	// Iterate through the string and count how many characters are vowels.
	return 0
}

func main() {
	fmt.Println(CountVowels("Hello World"))
	fmt.Println(CountVowels("Go is fun!"))
	fmt.Println(CountVowels(""))
	fmt.Println(CountVowels("AEIOUaeiou"))
}
