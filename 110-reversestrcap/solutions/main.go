package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseCap(s string) string {
	words := strings.Fields(s)
	var result []string
	for _, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			for i := 0; i < len(runes)-1; i++ {
				runes[i] = rune(strings.ToLower(string(runes[i]))[0])
			}
			runes[len(runes)-1] = rune(strings.ToUpper(string(runes[len(runes)-1]))[0])
			result = append(result, string(runes))
		}
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(reverseCap(os.Args[i]))
	}
}
