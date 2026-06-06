package main

import "fmt"

func ToUpper(s string) string {
	// TODO: Implement this function
	// Hint: Iterate through each character.
	// If it's a lowercase letter (between 'a' and 'z'),
	// convert it to uppercase by subtracting 32 from its ASCII value.
	// Otherwise, keep it as is.
	return ""
}

func main() {
	fmt.Println(ToUpper("Hello"))
	fmt.Println(ToUpper("Go is fun!"))
	fmt.Println(ToUpper("123abc"))
	fmt.Println(ToUpper(""))
}
