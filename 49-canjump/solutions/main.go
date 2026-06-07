package main

import "fmt"

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	// Each value represents exact steps to take forward
	pos := 0
	for pos < len(nums)-1 {
		steps := int(nums[pos])
		if steps == 0 {
			return false
		}
		pos += steps
		if pos >= len(nums)-1 {
			return true
		}
	}
	return pos == len(nums)-1
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
