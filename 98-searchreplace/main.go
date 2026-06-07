package main

import "fmt"

// SearchReplace replaces the first occurrence of old with new in text.
// If old is not found, return text unchanged.
func SearchReplace(text, old, new string) string {
	// TODO: Implement this function
	// Hint: Find the index of old in text.
	// If found, combine: text[:index] + new + text[index+len(old):]
	// If not found, return text as-is.
	return ""
}

func main() {
	fmt.Println(SearchReplace("Hello World", "World", "Go"))
	fmt.Println(SearchReplace("Hello World", "xyz", "Go"))
	fmt.Println(SearchReplace("Hello", "l", "L"))
	fmt.Println(SearchReplace("", "a", "b"))
}
