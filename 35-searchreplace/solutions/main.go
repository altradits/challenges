package main

import "fmt"

func SearchReplace(text, old, new string) string {
	index := FindSubstring(text, old)
	if index == -1 {
		return text
	}
	before := text[:index]
	after := text[index+len(old):]
	return before + new + after
}

func FindSubstring(text, substring string) int {
	if len(substring) == 0 {
		return 0
	}
	if len(substring) > len(text) {
		return -1
	}
	for i := 0; i <= len(text)-len(substring); i++ {
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
	fmt.Println(SearchReplace("Hello World", "World", "Go"))
	fmt.Println(SearchReplace("Hello World", "xyz", "Go"))
	fmt.Println(SearchReplace("Hello", "l", "L"))
	fmt.Println(SearchReplace("", "a", "b"))
}
