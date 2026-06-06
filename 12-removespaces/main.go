package main

import "fmt"

// RemoveSpaces returns a new string with all space characters removed.
func RemoveSpaces(s string) string {
	// TODO: Implement this function
	// Hint: Iterate through the string and build a new string
	// that only includes characters that are NOT spaces (' ').
	return ""
}

func main() {
	fmt.Println(RemoveSpaces("Hello World"))
	fmt.Println(RemoveSpaces("Go is fun!"))
	fmt.Println(RemoveSpaces(""))
	fmt.Println(RemoveSpaces("NoSpacesHere"))
}
