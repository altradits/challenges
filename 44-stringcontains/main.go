package main

import (
	"fmt"
	"strings"
)

func Contains(s, substr string) bool {
	if substr == "" {
		return true
	}
	return strings.Index(s, substr) != -1
}

func main() {
	fmt.Println(Contains("hello world", "world"))
	fmt.Println(Contains("hello world", "xyz"))
	fmt.Println(Contains("", "a"))
	fmt.Println(Contains("abc", ""))
}
