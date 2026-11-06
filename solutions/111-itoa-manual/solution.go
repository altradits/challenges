package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	var result []byte
	for n > 0 {
		digit := n % 10
		result = append(result, byte('0'+digit))
		n /= 10
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
