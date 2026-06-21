package main

import "fmt"

func FindSubstring(text, substring string) int {
	// Edge case: empty substring
	if len(substring) == 0 {
		return 0
	}

	// Edge case: substring longer than text
	if len(substring) > len(text) {
		return -1
	}

	// Loop through each possible starting position in text
	for i := 0; i <= len(text)-len(substring); i++ {
		// Check if substring matches at this position
		match := true
		for j := 0; j < len(substring); j++ {
			if text[i+j] != substring[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(FindSubstring("Hello World", "World"))
	fmt.Println(FindSubstring("Hello World", "hello"))
	fmt.Println(FindSubstring("Hello World", ""))
	fmt.Println(FindSubstring("Hello World", "xyz"))
	fmt.Println(FindSubstring("banana", "ana"))
}
