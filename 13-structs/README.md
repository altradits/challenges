# 28-structs — Structs in Go

## Challenge

Implement the following in `package piscine`:

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) Person
func (p Person) Greet() string    // "Hi, I'm Alice"
func (p Person) IsAdult() bool    // true if Age >= 18
```

## Examples

```
NewPerson("Alice", 30).Greet()   → "Hi, I'm Alice"
NewPerson("Bob", 15).IsAdult()   → false
NewPerson("Carol", 18).IsAdult() → true
```

Read `skills.md` before you start.
