package piscine

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func QuickSort(nums []int) {
	quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(nums, lo, hi)
	quickSortHelper(nums, lo, p-1)
	quickSortHelper(nums, p+1, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

func CountingSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	counts := make([]int, max+1)
	for _, v := range nums {
		counts[v]++
	}
	idx := 0
	for val, cnt := range counts {
		for i := 0; i < cnt; i++ {
			nums[idx] = val
			idx++
		}
	}
}
