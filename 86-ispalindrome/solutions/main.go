package main

import (
	"fmt"
)

func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	// Convert to lowercase and keep only alphanumeric
	var cleaned []rune
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			cleaned = append(cleaned, c)
		}
	}

	// Convert to lowercase
	for i := range cleaned {
		if cleaned[i] >= 'A' && cleaned[i] <= 'Z' {
			cleaned[i] = cleaned[i] - 'A' + 'a'
		}
	}

	// Check palindrome
	str := string(cleaned)
	return str == reverseString(str)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("Hello"))
	fmt.Println(IsPalindrome("A man a plan a canal Panama"))
	fmt.Println(IsPalindrome(""))
}
