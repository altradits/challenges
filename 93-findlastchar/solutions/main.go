package main

import "fmt"

func FindLastChar(s string, c rune) int {
	lastIndex := -1
	for i, ch := range s {
		if ch == c {
			lastIndex = i
		}
	}
	return lastIndex
}

func main() {
	fmt.Println(FindLastChar("Hello", 'l'))
	fmt.Println(FindLastChar("Hello", 'H'))
	fmt.Println(FindLastChar("Hello", 'z'))
	fmt.Println(FindLastChar("banana", 'a'))
	fmt.Println(FindLastChar("", 'a'))
}
