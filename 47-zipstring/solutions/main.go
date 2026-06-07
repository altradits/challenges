package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result []byte
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result = append(result, []byte(strconv.Itoa(count))...)
			result = append(result, s[i-1])
			count = 1
		}
	}
	// Add the last group
	result = append(result, []byte(strconv.Itoa(count))...)
	result = append(result, s[len(s)-1])

	return string(result)
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
