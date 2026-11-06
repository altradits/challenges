package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	if len(result) == 0 {
		return "\n"
	}
	return result + "\n"
}
