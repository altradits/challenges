package main

import (
	"fmt"
	"strings"
)

func Count(s, substr string) int {
	if substr == "" {
		return 0
	}
	count := 0
	start := 0
	for {
		idx := strings.Index(s[start:], substr)
		if idx == -1 {
			break
		}
		count++
		start += idx + 1 // +1 for overlapping
	}
	return count
}

func main() {
	fmt.Println(Count("hello hello", "l"))
	fmt.Println(Count("hello hello", "ll"))
	fmt.Println(Count("hello hello", "xyz"))
	fmt.Println(Count("", "a"))
}
