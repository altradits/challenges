package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	pos := 0
	for _, c := range arg {
		if pos < num {
			result += string(c)
		}
		pos++
		if pos == 2*num {
			pos = 0
		}
	}
	return result
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
