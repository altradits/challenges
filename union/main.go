package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := make(map[rune]bool)
	var result []rune
	for _, c := range s1 + s2 {
		if !seen[c] {
			seen[c] = true
			result = append(result, c)
		}
	}
	fmt.Println(string(result))
}
