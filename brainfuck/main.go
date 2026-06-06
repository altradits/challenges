package main

import (
	"fmt"
	"os"
)

func brainfuck(code string) {
	mem := make([]byte, 2048)
	ptr := 0
	pc := 0
	loops := []int{}
	for pc < len(code) {
		switch code[pc] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			fmt.Printf("%c", mem[ptr])
		case '[':
			if mem[ptr] == 0 {
				depth := 1
				for depth > 0 {
					pc++
					if code[pc] == '[' {
						depth++
					} else if code[pc] == ']' {
						depth--
					}
				}
			} else {
				loops = append(loops, pc)
			}
		case ']':
			if mem[ptr] != 0 {
				pc = loops[len(loops)-1]
			} else {
				loops = loops[:len(loops)-1]
			}
		}
		pc++
	}
	fmt.Println()
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	brainfuck(os.Args[1])
}
