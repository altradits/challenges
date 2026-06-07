package main

import (
	"fmt"
	"strings"
)

func CleanList(inputs []string) []string {
	var result []string
	for _, s := range inputs {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func main() {
	fmt.Println(CleanList([]string{"Hello", " World ", "Go"}))
	fmt.Println(CleanList([]string{"", "a", "", "b"}))
	fmt.Println(CleanList([]string{}))
	fmt.Println(CleanList([]string{"  "}))
}
