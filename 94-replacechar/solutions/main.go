package main

import "fmt"

func ReplaceChar(s string, old, new rune) string {
	result := ""
	for _, c := range s {
		if c == old {
			result += string(new)
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(ReplaceChar("Hello", 'l', 'L'))
	fmt.Println(ReplaceChar("Hello", 'z', 'X'))
	fmt.Println(ReplaceChar("", 'a', 'b'))
	fmt.Println(ReplaceChar("aaaa", 'a', 'b'))
}
