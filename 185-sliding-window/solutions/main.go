package piscine

func MaxSubarraySum(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	maxSum := windowSum
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func LongestUniqueSubstring(s string) int {
	charIndex := make(map[rune]int)
	maxLen := 0
	left := 0
	for right, ch := range s {
		if idx, ok := charIndex[ch]; ok && idx >= left {
			left = idx + 1
		}
		charIndex[ch] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func SmallestSubarrayWithSum(nums []int, target int) int {
	minLen := len(nums) + 1
	windowSum := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		windowSum += nums[right]
		for windowSum >= target {
			length := right - left + 1
			if length < minLen {
				minLen = length
			}
			windowSum -= nums[left]
			left++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func AnagramCount(s, pattern string) int {
	if len(pattern) > len(s) {
		return 0
	}
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(pattern); i++ {
		need[pattern[i]]++
		window[s[i]]++
	}
	count := 0
	if mapsEqual(window, need) {
		count++
	}
	for i := len(pattern); i < len(s); i++ {
		window[s[i]]++
		old := s[i-len(pattern)]
		window[old]--
		if window[old] == 0 {
			delete(window, old)
		}
		if mapsEqual(window, need) {
			count++
		}
	}
	return count
}

func mapsEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
