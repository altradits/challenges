package main

import "fmt"

func FindChar(s string, c rune) int {
	for i, ch := range s {
		if ch == c {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FindChar("Hello", 'l'))
	fmt.Println(FindChar("Hello", 'H'))
	fmt.Println(FindChar("Hello", 'z'))
	fmt.Println(FindChar("", 'a'))
}
