package main

import "fmt"

func ToUpper(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(c - 32)
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(ToUpper("Hello"))
	fmt.Println(ToUpper("Go is fun!"))
	fmt.Println(ToUpper("123abc"))
	fmt.Println(ToUpper(""))
}
