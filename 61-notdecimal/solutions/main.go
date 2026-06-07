package main

import (
	"fmt"
	"strings"
)

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}

	// Check if it's a valid number with a decimal point
	hasDecimal := false
	hasNonZeroAfterDecimal := false
	dotIndex := -1

	for i, c := range dec {
		if c == '.' {
			if hasDecimal {
				return dec + "\n" // multiple dots, not valid
			}
			hasDecimal = true
			dotIndex = i
		} else if c == '-' && i == 0 {
			// allow negative sign at start
		} else if c < '0' || c > '9' {
			return dec + "\n" // not a number
		}
	}

	if !hasDecimal {
		return dec + "\n" // no decimal point
	}

	// Check if there's only zero after the decimal
	afterDot := dec[dotIndex+1:]
	for _, c := range afterDot {
		if c != '0' {
			hasNonZeroAfterDecimal = true
			break
		}
	}

	if !hasNonZeroAfterDecimal {
		return dec + "\n" // only zeros after decimal
	}

	// Remove the decimal point
	result := strings.Replace(dec, ".", "", 1)
	return result + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
