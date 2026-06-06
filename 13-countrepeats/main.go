package main

import "fmt"

// CountRepeats returns the number of groups of consecutive repeated characters.
// A "repeat" is when the same character appears consecutively two or more times.
// Count each group as one.
func CountRepeats(s string) int {
	// TODO: Implement this function
	// Hint: Track the previous character and whether you're currently
	// inside a repeat group. Only count when entering a new repeat group.
	return 0
}

func main() {
	fmt.Println(CountRepeats("hello"))
	fmt.Println(CountRepeats("helloo"))
	fmt.Println(CountRepeats("heelloo"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats(""))
	fmt.Println(CountRepeats("aaa"))
}
