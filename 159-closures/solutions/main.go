package piscine

func MakeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func MakeAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func Filter(nums []int, keep func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
