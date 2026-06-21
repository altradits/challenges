package main

import "fmt"

func ToLower(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			result += string(c + 32)
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(ToLower("Hello"))
	fmt.Println(ToLower("Go is fun!"))
	fmt.Println(ToLower("123ABC"))
	fmt.Println(ToLower(""))
}
