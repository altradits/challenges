package piscine

func TwoSum(nums []int, target int) (int, int) {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			return lo, hi
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return -1, -1
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func ReverseInPlace(nums []int) {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}

func ContainerWithMostWater(heights []int) int {
	lo, hi := 0, len(heights)-1
	maxWater := 0
	for lo < hi {
		h := heights[lo]
		if heights[hi] < h {
			h = heights[hi]
		}
		water := h * (hi - lo)
		if water > maxWater {
			maxWater = water
		}
		if heights[lo] < heights[hi] {
			lo++
		} else {
			hi--
		}
	}
	return maxWater
}
