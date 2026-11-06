package main

import "fmt"

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		if i > 0 && i%4 == 0 {
			fmt.Println()
		} else if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%02x", b)
	}
	fmt.Println()
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
