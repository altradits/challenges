package piscine

// RevConcatAlternate reverses both slices, emits the longer slice's extra
// head, then strictly alternates starting with slice1.
func RevConcatAlternate(slice1, slice2 []int) []int {
	rev1 := make([]int, len(slice1))
	for i := range slice1 {
		rev1[i] = slice1[len(slice1)-1-i]
	}
	rev2 := make([]int, len(slice2))
	for i := range slice2 {
		rev2[i] = slice2[len(slice2)-1-i]
	}
	var result []int
	if len(rev1) >= len(rev2) {
		d := len(rev1) - len(rev2)
		result = append(result, rev1[:d]...)
		for i := 0; i < len(rev2); i++ {
			result = append(result, rev1[d+i], rev2[i])
		}
	} else {
		d := len(rev2) - len(rev1)
		result = append(result, rev2[:d]...)
		for i := 0; i < len(rev1); i++ {
			result = append(result, rev1[i], rev2[d+i])
		}
	}
	return result
}
