package main

import "fmt"

// CountWordsAdvanced counts words where words are sequences of letters and digits.
// Punctuation and spaces separate words.
func CountWordsAdvanced(s string) int {
	// TODO: Implement this function
	// Hint: Track whether you're inside a word.
	// A word character is a letter (a-z, A-Z) or digit (0-9).
	return 0
}

func main() {
	fmt.Println(CountWordsAdvanced("Hello, World!"))
	fmt.Println(CountWordsAdvanced("Go is fun"))
	fmt.Println(CountWordsAdvanced("hello2world"))
	fmt.Println(CountWordsAdvanced(""))
}
