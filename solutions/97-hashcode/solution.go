package piscine

func HashCode(dec string) string {
	size := len(dec)
	result := make([]byte, size)

	for i, c := range dec {
		hashed := (int(c) + size) % 127
		if hashed < 32 {
			hashed += 33
		}
		result[i] = byte(hashed)
	}

	return string(result)
}
