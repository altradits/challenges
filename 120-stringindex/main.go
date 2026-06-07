package main

import (
	"fmt"
)

func Index(s, substr string) int {
	if substr == "" {
		return 0
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("hello world", "world"))
	fmt.Println(Index("hello world", "xyz"))
	fmt.Println(Index("abcabc", "b"))
	fmt.Println(Index("", "a"))
}
