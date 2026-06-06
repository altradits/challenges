package main

import (
	"fmt"
)

func HasSuffix(s, suffix string) bool {
	if suffix == "" {
		return true
	}
	if len(s) < len(suffix) {
		return false
	}
	return s[len(s)-len(suffix):] == suffix
}

func main() {
	fmt.Println(HasSuffix("hello world", "world"))
	fmt.Println(HasSuffix("hello world", "hello"))
	fmt.Println(HasSuffix("", ""))
	fmt.Println(HasSuffix("abc", "abcd"))
}
