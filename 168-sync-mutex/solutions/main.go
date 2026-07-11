package piscine

import "sync"

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += n
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *SafeCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count = 0
}

type SafeMap struct {
	mu    sync.RWMutex
	store map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{store: make(map[string]int)}
}

func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = value
}

func (m *SafeMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.store[key]
	return v, ok
}

func (m *SafeMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}

func (m *SafeMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.store)
}
