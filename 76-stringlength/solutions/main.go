package main

import "fmt"

func StringLength(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {
	fmt.Println(StringLength("Hello"))
	fmt.Println(StringLength(""))
	fmt.Println(StringLength("Go is fun"))
	fmt.Println(StringLength("a"))
}
