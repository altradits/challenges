package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evalRPN(tokens []string) (int, bool) {
	var stack []int
	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" || t == "%" {
			if len(stack) < 2 {
				return 0, false
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch t {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					return 0, false
				}
				res = a / b
			case "%":
				if b == 0 {
					return 0, false
				}
				res = a % b
			}
			stack = append(stack, res)
		} else {
			n, err := strconv.Atoi(t)
			if err != nil {
				return 0, false
			}
			stack = append(stack, n)
		}
	}
	if len(stack) != 1 {
		return 0, false
	}
	return stack[0], true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	tokens := strings.Fields(os.Args[1])
	if len(tokens) == 0 {
		fmt.Println("Error")
		return
	}
	result, ok := evalRPN(tokens)
	if !ok {
		fmt.Println("Error")
		return
	}
	fmt.Println(result)
}
