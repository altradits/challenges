package main

import (
	"fmt"
	"os"
)

func isBracketed(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		if isBracketed(os.Args[i]) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
