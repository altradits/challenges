package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	for _, arg := range os.Args[1:] {
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}
	}
	fmt.Println("Invalid Option")
}
