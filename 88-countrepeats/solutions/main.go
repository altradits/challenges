package main

import "fmt"

func CountRepeats(s string) int {
	if len(s) < 2 {
		return 0
	}

	count := 0
	inRepeat := false

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if !inRepeat {
				count++
				inRepeat = true
			}
		} else {
			inRepeat = false
		}
	}

	return count
}

func main() {
	fmt.Println(CountRepeats("hello"))
	fmt.Println(CountRepeats("helloo"))
	fmt.Println(CountRepeats("heelloo"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats(""))
	fmt.Println(CountRepeats("aaa"))
}
