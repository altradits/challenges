package main

import "fmt"

// ToLower converts all uppercase letters in a string to lowercase.
// Non-alphabetic characters remain unchanged.
func ToLower(s string) string {
	// TODO: Implement this function
	// Hint: Iterate through each character.
	// If it's an uppercase letter (between 'A' and 'Z'),
	// convert it to lowercase by adding 32 to its ASCII value.
	// Otherwise, keep it as is.
	return ""
}

func main() {
	fmt.Println(ToLower("Hello"))
	fmt.Println(ToLower("Go is fun!"))
	fmt.Println(ToLower("123ABC"))
	fmt.Println(ToLower(""))
}
