package main

import "fmt"

func WordCount(s string) int {
	count := 0
	inWord := false

	for _, c := range s {
		if c == ' ' {
			inWord = false
		} else if !inWord {
			count++
			inWord = true
		}
	}

	return count
}

func main() {
	fmt.Println(WordCount("Hello World"))
	fmt.Println(WordCount("Go is fun!"))
	fmt.Println(WordCount(""))
	fmt.Println(WordCount("   "))
	fmt.Println(WordCount("One  two   three"))
}
