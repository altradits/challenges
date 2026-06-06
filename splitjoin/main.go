package main

import "fmt"

func Split(s string, sep string) []string {
	// TODO: Implement this function
	// Hint: Iterate through the string and find each occurrence of sep.
	// Build substrings between separators and add them to a slice.
	return nil
}

func Join(arr []string, sep string) string {
	// TODO: Implement this function
	// Hint: Iterate through the slice and concatenate elements with sep between them.
	return ""
}

func main() {
	// Test Split
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("Hello World", " "))
	fmt.Printf("%q\n", Split("abc", ""))
	fmt.Printf("%q\n", Split("", ","))

	// Test Join
	fmt.Println(Join([]string{"a", "b", "c"}, ","))
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]string{"a", "b", "c"}, ""))
	fmt.Println(Join([]string{""}, ","))
}
