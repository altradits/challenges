package main

import (
	"fmt"
	"strings"
)

func Join(words []string, sep string) string {
	if len(words) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(words[0])
	for _, w := range words[1:] {
		b.WriteString(sep)
		b.WriteString(w)
	}
	return b.String()
}

func main() {
	fmt.Printf("%q\n", Join([]string{"a", "b", "c"}, "-"))
	fmt.Printf("%q\n", Join([]string{"hello", "world"}, " "))
	fmt.Printf("%q\n", Join([]string{}, ","))
	fmt.Printf("%q\n", Join([]string{"single"}, "|"))
}
