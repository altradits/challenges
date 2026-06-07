package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	result := make([]int, len(slice1)+len(slice2))
	copy(result, slice1)
	copy(result[len(slice1):], slice2)
	return result
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
