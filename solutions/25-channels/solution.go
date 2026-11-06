package piscine

import "sync"

func Generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func Merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	forward := func(in <-chan int) {
		defer wg.Done()
		for v := range in {
			out <- v
		}
	}
	go forward(a)
	go forward(b)
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func WorkerPool(jobs <-chan int, numWorkers int, fn func(int) int) <-chan int {
	if numWorkers <= 0 {
		numWorkers = 1
	}
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				out <- fn(job)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
