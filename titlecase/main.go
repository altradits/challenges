package main

import "fmt"

func TitleCase(s string) string {
	// TODO: Implement this function
	// Hint: Track whether the next character should be uppercase.
	// 1. Iterate through the string.
	// 2. If the current character is a letter and the previous character was a space (or it's the first character), make it uppercase.
	// 3. Otherwise, make it lowercase.
	// 4. Keep non-letter characters unchanged.
	return ""
}

func main() {
	fmt.Println(TitleCase("hello world"))
	fmt.Println(TitleCase("Go is fun!"))
	fmt.Println(TitleCase(""))
	fmt.Println(TitleCase("multiple   spaces   here"))
	fmt.Println(TitleCase("ALL CAPS"))
}
