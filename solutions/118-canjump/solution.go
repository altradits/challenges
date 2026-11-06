package piscine

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
