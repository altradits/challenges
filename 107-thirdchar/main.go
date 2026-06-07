package main

import "fmt"

func ThirdChar(s string) string {
	// TODO: Implement this function
	// Hint: Use for...range with index.
	// Check if (index + 1) % 3 == 0 to get every 3rd character.
	return ""
}

func main() {
	fmt.Print(ThirdChar("123456789"))
	fmt.Print(ThirdChar(""))
	fmt.Print(ThirdChar("a b c d e f"))
	fmt.Print(ThirdChar("12"))
}
