package main

import "fmt"

func FirstChar(s string) string {
	if s == "" {
		return ""
	}
	for _, c := range s {
		return string(c)
	}
	return ""
}

func main() {
	fmt.Println(FirstChar("Hello"))
	fmt.Println(FirstChar(""))
	fmt.Println(FirstChar("G"))
	fmt.Println(FirstChar("Go is fun"))
}
