package main

import "fmt"

// PrintIfNot returns "G\n" if the string length is less than 3, otherwise "Invalid Input\n".
func PrintIfNot(str string) string {
	// TODO: Implement this function
	// Hint: This is the inverse of PrintIf.
	// Check if len(str) < 3
	return ""
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
