# 29-pointers — Pointers in Go

## Challenge

Implement in `package piscine`:

```go
func Double(n *int)                      // doubles the int at the pointer in-place
func Swap(a, b *int)                     // swaps values at the two pointers
func NewPerson(name string, age int) *Person  // returns a pointer to a new Person
```

With `Person` defined as:
```go
type Person struct{ Name string; Age int }
```

## Examples

```go
x := 5
Double(&x)   // x is now 10

a, b := 3, 7
Swap(&a, &b) // a is 7, b is 3

p := NewPerson("Alice", 30)
fmt.Println(p.Name)  // Alice  (p is *Person)
```

Read `skills.md` before you start.
