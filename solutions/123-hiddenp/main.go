package main

import (
	"fmt"
	"os"
)

func isHidden(s1, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	if isHidden(os.Args[1], os.Args[2]) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
