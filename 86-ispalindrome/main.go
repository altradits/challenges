package main

import "fmt"

// IsPalindrome returns true if the string is a palindrome, false otherwise.
// Ignore case and non-alphanumeric characters.
func IsPalindrome(s string) bool {
	// TODO: Implement this function
	// Hint 1: Convert the string to lowercase first.
	// Hint 2: Build a new string containing only alphanumeric characters.
	// Hint 3: Compare the cleaned string with its reverse.
	return false
}

func main() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("Hello"))
	fmt.Println(IsPalindrome("A man a plan a canal Panama"))
	fmt.Println(IsPalindrome(""))
}
