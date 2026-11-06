package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	if len(slice1) >= len(slice2) {
		// slice1 is larger or equal, start with slice1
		for i := 0; i < len(slice2); i++ {
			result = append(result, slice1[i], slice2[i])
		}
		result = append(result, slice1[len(slice2):]...)
	} else {
		// slice2 is larger, start with slice2
		for i := 0; i < len(slice1); i++ {
			result = append(result, slice2[i], slice1[i])
		}
		result = append(result, slice2[len(slice1):]...)
	}

	return result
}
