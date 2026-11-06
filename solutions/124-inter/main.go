package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := make(map[rune]bool)
	for _, c := range s1 {
		if !seen[c] {
			for _, d := range s2 {
				if c == d {
					fmt.Printf("%c", c)
					seen[c] = true
					break
				}
			}
		}
	}
	fmt.Println()
}
