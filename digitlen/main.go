package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	// Check if base is valid (between 2 and 36)
	if base < 2 || base > 36 {
		return -1
	}

	// If n is negative, reverse the sign
	if n < 0 {
		n = -n
	}

	// Count the number of digits
	count := 0
	for n > 0 {
		n = n / base
		count++
	}

	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}