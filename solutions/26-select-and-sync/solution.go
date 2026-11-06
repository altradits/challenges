package piscine

import (
	"errors"
	"sync"
	"time"
)

var ErrTimeout = errors.New("operation timed out")

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func Timeout(fn func() string, d time.Duration) (string, error) {
	done := make(chan string, 1)
	go func() {
		done <- fn()
	}()
	select {
	case v := <-done:
		return v, nil
	case <-time.After(d):
		return "", ErrTimeout
	}
}

func NewOnce(fn func() string) func() string {
	var once sync.Once
	var result string
	return func() string {
		once.Do(func() {
			result = fn()
		})
		return result
	}
}
