package main

import (
	"fmt"
)

func HasPrefix(s, prefix string) bool {
	if prefix == "" {
		return true
	}
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func main() {
	fmt.Println(HasPrefix("hello world", "hello"))
	fmt.Println(HasPrefix("hello world", "world"))
	fmt.Println(HasPrefix("", ""))
	fmt.Println(HasPrefix("abc", "abcd"))
}
