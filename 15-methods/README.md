# 15-methods — Value vs Pointer Receivers

## Task

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

---

## Time limit

**15 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 15
```

or run the tests directly:

```bash
go test ./15-methods/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 15` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
