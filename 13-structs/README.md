# 13-structs — Structs in Go

## Task

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

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 13
```

or run the tests directly:

```bash
go test ./13-structs/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 13` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
