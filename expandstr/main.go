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
	if len(words) == 0 {
		return
	}
	fmt.Println(strings.Join(words, "   "))
}
