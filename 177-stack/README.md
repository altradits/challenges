# 177-stack — Stack (LIFO)

## Task

Implement a generic integer stack in `package piscine`:

```go
type Stack struct { /* unexported fields */ }

func NewStack() *Stack
func (s *Stack) Push(val int)
func (s *Stack) Pop() (int, bool)   // returns false if empty
func (s *Stack) Peek() (int, bool)  // look without removing
func (s *Stack) IsEmpty() bool
func (s *Stack) Size() int

// IsBalanced checks if the brackets in s are correctly balanced.
// Brackets: () [] {}
// Example: "({[]})" → true, "([)]" → false, "(((" → false
func IsBalanced(s string) bool
```

## Examples

```
st := NewStack()
st.Push(1)
st.Push(2)
st.Push(3)
st.Peek()    →  3, true
st.Pop()     →  3, true
st.Pop()     →  2, true
st.Size()    →  1

IsBalanced("({[]})")  →  true
IsBalanced("([)]")    →  false
IsBalanced("hello")   →  true  (no brackets = balanced)
IsBalanced("((")      →  false
```

Read `skills.md` before you start.

---

## Time limit

**30 minutes** (round 1). Beat it on every repeat — each spaced round targets a shorter time as the solution moves into long-term memory.

## Blind rule

No docs, no search, no AI, no copying. Close `skills.md`, close the reference solution, turn off autocomplete if you can. Write every line from memory — that is the whole point.

## How to check

From the repo root:

```bash
go run ./cmd/drill check 177
```

or run the tests directly:

```bash
go test ./177-stack/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 177` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
