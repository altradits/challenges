package piscine

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Add(n int) {
	c.count += n
}

func (c *Counter) Reset() {
	c.count = 0
}

func (c Counter) Value() int {
	return c.count
}
