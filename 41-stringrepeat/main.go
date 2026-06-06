package main

import (
	"fmt"
	"strings"
)

func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	fmt.Printf("%q\n", Repeat("abc", 3))
	fmt.Printf("%q\n", Repeat("x", 5))
	fmt.Printf("%q\n", Repeat("go", 0))
	fmt.Printf("%q\n", Repeat("", 10))
}
