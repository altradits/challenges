package main

import (
	"fmt"
)

func Max(a, b rune) rune {
	if a > b {
		return a
	}
	return b
}

func Min(a, b rune) rune {
	if a < b {
		return a
	}
	return b
}

func Reduce(s string, f func(rune, rune) rune) rune {
	if len(s) == 0 {
		return 0
	}
	var result rune
	first := true
	for _, r := range s {
		if first {
			result = r
			first = false
		} else {
			result = f(result, r)
		}
	}
	return result
}

func main() {
	fmt.Printf("%q\n", Reduce("abc", Max))
	fmt.Printf("%q\n", Reduce("hello", Min))
	fmt.Printf("%q\n", Reduce("123", func(a, b rune) rune { return a + b }))
	fmt.Printf("%q\n", Reduce("", Max))
}
