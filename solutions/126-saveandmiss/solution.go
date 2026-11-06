package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	pos := 0
	for _, c := range arg {
		if pos < num {
			result += string(c)
		}
		pos++
		if pos == 2*num {
			pos = 0
		}
	}
	return result
}
