package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}

	if len(slice) == 0 {
		fmt.Println([]int{})
		return
	}

	var result [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	fmt.Println(result)
}
