package main

import (
	"fmt"
	"os"
	"strconv"
)

func toRoman(n int) string {
	if n <= 0 || n >= 4000 {
		return "ERROR: cannot convert to roman digit"
	}
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string
	for i, v := range vals {
		for n >= v {
			n -= v
			result += syms[i]
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	fmt.Println(toRoman(n))
}
