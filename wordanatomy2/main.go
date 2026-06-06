package main

import "fmt"

// Wordanatomy matches the exact signature you requested
func Wordanatomy(word string, pref []string, suff []string) string {
	matchedPrefix := ""
	matchedSuffix := ""

	// 1. Find the first matching prefix
	for _, p := range pref {
		if len(word) >= len(p) && word[:len(p)] == p {
			matchedPrefix = p
			break // Stop at the first match
		}
	}

	// 2. Find the first matching suffix
	for _, s := range suff {
		if len(word) >= len(s) && word[len(word)-len(s):] == s {
			matchedSuffix = s
			break // Stop at the first match
		}
	}

	// 3. Format the exact output string based on your requirements
	// Using %q automatically wraps the string variables in double quotes ""
	return fmt.Sprintf("prefix: %q, suffix: %q", matchedPrefix, matchedSuffix)
}

func main() {
	prefixes := []string{"an", "en", "un"}
	suffixes := []string{"ible", "able"}

	// Test Case 1: Both Prefix and Suffix exist
	fmt.Println(Wordanatomy("understandable", prefixes, suffixes))
	// Output: prefix: "un", suffix: "able"

	// Test Case 2: Only Suffix exists
	fmt.Println(Wordanatomy("understand", prefixes, suffixes))
	// Output: prefix: "un", suffix: ""

	// Test Case 3: Neither exists
	fmt.Println(Wordanatomy("apple", prefixes, suffixes))
	// Output: prefix: "", suffix: ""
}