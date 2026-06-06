package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	if sep == "" {
		result := make([]string, len(s))
		for i, r := range s {
			result[i] = string(r)
		}
		return result
	}
	var result []string
	start := 0
	for {
		idx := strings.Index(s[start:], sep)
		if idx == -1 {
			result = append(result, s[start:])
			break
		}
		result = append(result, s[start:start+idx])
		start += idx + len(sep)
	}
	return result
}

func main() {
	fmt.Printf("%q\n", Split("a,b,c", ","))
	fmt.Printf("%q\n", Split("a:b:c", ":"))
	fmt.Printf("%q\n", Split("abc", "x"))
	fmt.Printf("%q\n", Split("", ","))
}
