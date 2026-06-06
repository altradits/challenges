package main

import "fmt"

// CountChar returns the number of times c appears in str.
func CountChar(str string, c rune) int {
	// TODO: Implement this function
	// Hint: Initialize a counter to 0.
	// Iterate through the string and increment when you find a match.
	// Return the final count.
	return 0
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", ' '))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
