package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := strings.Fields(os.Args[1])
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
