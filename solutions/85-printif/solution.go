package piscine

func PrintIf(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
