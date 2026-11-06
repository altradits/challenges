package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	words := strings.Fields(os.Args[1])
	if len(words) == 0 {
		fmt.Println()
		return
	}
	fmt.Println(strings.Join(words, " "))
}
