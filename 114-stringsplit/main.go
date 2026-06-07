package main

import (
	"fmt"
	"strings"
)

func SplitWords(s string) []string {
	fields := strings.Fields(s)
	result := make([]string, len(fields))
	copy(result, fields)
	return result
}

func main() {
	fmt.Printf("%q\n", SplitWords("Hello World"))
	fmt.Printf("%q\n", SplitWords("  Go   is  fun  "))
	fmt.Printf("%q\n", SplitWords(""))
	fmt.Printf("%q\n", SplitWords("single"))
}
