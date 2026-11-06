package piscine

func BinarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func SearchInsertPosition(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func FindFirstOccurrence(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			result = mid
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func FindLastOccurrence(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	result := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			result = mid
			lo = mid + 1
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}
