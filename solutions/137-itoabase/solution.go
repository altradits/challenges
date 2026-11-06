package piscine

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	negative := false
	if value < 0 {
		negative = true
		value = -value
	}

	digits := "0123456789ABCDEF"
	var result []byte

	for value > 0 {
		digit := value % base
		result = append(result, digits[digit])
		value /= base
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	if negative {
		return "-" + string(result)
	}
	return string(result)
}
