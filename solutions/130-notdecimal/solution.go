package piscine

import (
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

	// Remove the decimal point, then strip leading zeros ("0.1" -> "1").
	result := strings.Replace(dec, ".", "", 1)
	neg := strings.HasPrefix(result, "-")
	digits := strings.TrimPrefix(result, "-")
	digits = strings.TrimLeft(digits, "0")
	if digits == "" {
		digits = "0"
	}
	if neg {
		digits = "-" + digits
	}
	return digits + "\n"
}
