package main

import "fmt"

// PrintIf returns "G\n" if the string length is 0 or >= 3, otherwise "Invalid Input\n".
func PrintIf(str string) string {
	// TODO: Implement this function
	// Hint: Check if len(str) == 0 || len(str) >= 3
	return ""
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
