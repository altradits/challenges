package piscine

func RetainFirstHalf(s string) string {
	if len(s) <= 1 {
		return s
	}
	half := len(s) / 2
	return s[:half]
}
