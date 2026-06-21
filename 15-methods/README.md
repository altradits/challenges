# 30-methods — Value vs Pointer Receivers

## Challenge

Given:
```go
type Counter struct{ count int }
```

Implement in `package piscine`:
```go
func (c *Counter) Increment()       // adds 1 to count
func (c *Counter) Add(n int)        // adds n to count
func (c *Counter) Reset()           // sets count to 0
func (c Counter)  Value() int       // returns current count (read-only — value receiver)
```

## Examples

```go
c := Counter{}
c.Increment()   // count = 1
c.Add(4)        // count = 5
fmt.Println(c.Value())  // 5
c.Reset()
fmt.Println(c.Value())  // 0
```

Read `skills.md` before you start.
