package main

import "fmt"

func FindSubString(text, substring string) int {
	
	for i, word := range text {

		if len(substring) == 0 {
			return 0
		}

		

		if word == substring {
			return i
		}
	}
}

func main() {
	fmt.Println(FindSubString("Hello World", "World"))
	fmt.Println(FindSubString("Hello World", "hello"))
	fmt.Println(FindSubString("Hello World", ""))
	fmt.Println(FindSubString("Hello World", "xyz"))
	fmt.Println(FindSubString("banana", "ana"))
}