package piscine

import "fmt"

func Describe(v any) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("integer: %d", val)
	case string:
		return fmt.Sprintf("string: %s", val)
	case bool:
		return fmt.Sprintf("bool: %t", val)
	case []int:
		return fmt.Sprintf("int slice of length %d", len(val))
	default:
		return "unknown type"
	}
}

func Sum(values []any) float64 {
	total := 0.0
	for _, v := range values {
		switch n := v.(type) {
		case int:
			total += float64(n)
		case float64:
			total += n
		}
	}
	return total
}
