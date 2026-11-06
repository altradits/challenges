package piscine

func ReplaceAll(text, old, new string) string {
	if len(old) == 0 {
		return text
	}

	var result []byte
	i := 0
	for i < len(text) {
		if i <= len(text)-len(old) && text[i:i+len(old)] == old {
			result = append(result, new...)
			i += len(old)
		} else {
			result = append(result, text[i])
			i++
		}
	}
	return string(result)
}
