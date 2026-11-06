package piscine

import "sync"

type poolTask struct {
	index int
	job   func() int
}

func RunPool(jobs []func() int, workers int) []int {
	if workers < 1 {
		workers = 1
	}
	results := make([]int, len(jobs))
	jobsCh := make(chan poolTask)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range jobsCh {
				results[t.index] = t.job()
			}
		}()
	}
	for i, j := range jobs {
		jobsCh <- poolTask{index: i, job: j}
	}
	close(jobsCh)
	wg.Wait()
	return results
}
