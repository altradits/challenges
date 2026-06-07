package main

import "fmt"

func CountVowels(s string) int {
	count := 0
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountVowels("Hello World"))
	fmt.Println(CountVowels("Go is fun!"))
	fmt.Println(CountVowels(""))
	fmt.Println(CountVowels("AEIOUaeiou"))
}
