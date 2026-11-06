package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanVals = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanSyms = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// toRoman converts n (1-3999) to a Roman numeral.
func toRoman(n int) string {
	if n <= 0 || n >= 4000 {
		return "ERROR: cannot convert to roman digit"
	}
	var result string
	for i, v := range romanVals {
		for n >= v {
			n -= v
			result += romanSyms[i]
		}
	}
	return result
}

// decomposition expands n into additive parts, with subtractive pairs in
// parentheses: 999 -> "(M-C)+(C-X)+(X-I)".
func decomposition(n int) string {
	var parts []string
	for i, v := range romanVals {
		for n >= v {
			n -= v
			sym := romanSyms[i]
			if len(sym) == 2 {
				parts = append(parts, "("+string(sym[1])+"-"+string(sym[0])+")")
			} else {
				parts = append(parts, sym)
			}
		}
	}
	return strings.Join(parts, "+")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 || n >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}
	fmt.Println(decomposition(n))
	fmt.Println(toRoman(n))
}
