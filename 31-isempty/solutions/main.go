package main

import "fmt"

func IsEmpty(s string) bool {
	for range s {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsEmpty("Hello"))
	fmt.Println(IsEmpty(""))
	fmt.Println(IsEmpty(" "))
	fmt.Println(IsEmpty("G"))
}
