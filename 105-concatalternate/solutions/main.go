package main

import "fmt"

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

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
