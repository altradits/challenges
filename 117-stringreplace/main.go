package main

import (
	"fmt"
	"strings"
)

func Replace(s, old, new string) string {
	if old == "" {
		return s
	}
	var b strings.Builder
	start := 0
	for {
		idx := strings.Index(s[start:], old)
		if idx == -1 {
			break
		}
		b.WriteString(s[start : start+idx])
		b.WriteString(new)
		start += idx + len(old)
	}
	b.WriteString(s[start:])
	return b.String()
}

func main() {
	fmt.Printf("%q\n", Replace("hello world", "l", "L"))
	fmt.Printf("%q\n", Replace("foo bar foo", "foo", "baz"))
	fmt.Printf("%q\n", Replace("abc", "x", "y"))
	fmt.Printf("%q\n", Replace("", "a", "b"))
}
