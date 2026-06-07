package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	seen1 := make(map[rune]bool)
	seen2 := make(map[rune]bool)
	seenBoth := make(map[rune]bool)

	for _, c := range str1 {
		seen1[c] = true
	}
	for _, c := range str2 {
		seen2[c] = true
	}

	count := 0
	for c := range seen1 {
		if !seen2[c] && !seenBoth[c] {
			count++
			seenBoth[c] = true
		}
	}
	for c := range seen2 {
		if !seen1[c] && !seenBoth[c] {
			count++
			seenBoth[c] = true
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
