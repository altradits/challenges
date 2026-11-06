package piscine

func RetainFirstHalf(str string) string {
	if len(str) <= 1 {
		return str
	}
	half := len(str) / 2
	return str[:half]
}
