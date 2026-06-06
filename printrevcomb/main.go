package main

import "fmt"

func main() {
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {
				fmt.Printf("%d%d%d", a, b, c)
				if a > 2 || b > 1 || c > 0 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}
