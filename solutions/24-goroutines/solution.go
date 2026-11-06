package piscine

import "sync"

func ConcurrentMap(data []int, fn func(int) int) []int {
	out := make([]int, len(data))
	var wg sync.WaitGroup
	for i, v := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out[i] = fn(v)
		}()
	}
	wg.Wait()
	return out
}

func ParallelSum(data []int, numWorkers int) int {
	if len(data) == 0 {
		return 0
	}
	if numWorkers <= 0 {
		numWorkers = 1
	}
	if numWorkers > len(data) {
		numWorkers = len(data)
	}
	chunk := (len(data) + numWorkers - 1) / numWorkers
	sums := make([]int, numWorkers)
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := w * chunk
			end := start + chunk
			if end > len(data) {
				end = len(data)
			}
			for _, v := range data[start:end] {
				sums[w] += v
			}
		}()
	}
	wg.Wait()
	total := 0
	for _, s := range sums {
		total += s
	}
	return total
}
