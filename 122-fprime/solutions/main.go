package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		return
	}
	first := true
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			first = false
			n /= i
		}
	}
	fmt.Println()
}
