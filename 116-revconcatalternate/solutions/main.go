package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	// Reverse both slices
	rev1 := make([]int, len(slice1))
	for i := 0; i < len(slice1); i++ {
		rev1[i] = slice1[len(slice1)-1-i]
	}

	rev2 := make([]int, len(slice2))
	for i := 0; i < len(slice2); i++ {
		rev2[i] = slice2[len(slice2)-1-i]
	}

	if len(slice1) >= len(slice2) {
		// slice1 is larger or equal, start with slice1
		for i := 0; i < len(slice2); i++ {
			result = append(result, rev1[i], rev2[i])
		}
		result = append(result, rev1[len(slice2):]...)
	} else {
		// slice2 is larger, start with slice2
		for i := 0; i < len(slice1); i++ {
			result = append(result, rev2[i], rev1[i])
		}
		result = append(result, rev2[len(slice1):]...)
	}

	return result
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
