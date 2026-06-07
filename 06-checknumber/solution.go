package main

import "fmt"

func CheckNumber(arg string) bool {
	
	for _, char := range arg {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
