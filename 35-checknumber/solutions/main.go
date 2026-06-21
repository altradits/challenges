package main

import "fmt"

func CheckNumber(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
	fmt.Println(CheckNumber("123"))
	fmt.Println(CheckNumber(""))
}
