package piscine

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
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n = n / base
		count++
	}

	return count
}
