package main

import "fmt"

// Join concatenates a slice of strings with a separator between them.
func Join(elems []string, sep string) string {
	// TODO: Implement this function
	// Hint: Handle the empty slice case first.
	// Start with the first element, then add sep + element for each remaining.
	return ""
}

func main() {
	fmt.Println(Join([]string{"a", "b", "c"}, ","))
	fmt.Println(Join([]string{"Hello", "World"}, " "))
	fmt.Println(Join([]string{"a", "b", "c"}, ""))
	fmt.Println(Join([]string{""}, ","))
	fmt.Println(Join([]string{}, ","))
}
