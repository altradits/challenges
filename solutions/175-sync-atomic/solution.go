package piscine

import "sync/atomic"

type Counter struct {
	v atomic.Int64
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Add(n int64) int64 {
	return c.v.Add(n)
}

func (c *Counter) Value() int64 {
	return c.v.Load()
}

func CASAdd(addr *int64, delta int64) int64 {
	for {
		old := atomic.LoadInt64(addr)
		if atomic.CompareAndSwapInt64(addr, old, old+delta) {
			return old + delta
		}
	}
}
