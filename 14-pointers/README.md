# 14-pointers — Pointers in Go

## Task

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

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 14
```

or run the tests directly:

```bash
go test ./14-pointers/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 14` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
