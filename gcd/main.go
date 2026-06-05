package main

import "fmt"

func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
