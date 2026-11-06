# 176-linked-list — Singly Linked List

## Challenge

Implement a singly linked list in `package piscine`:

```go
type Node struct {
    Val  int
    Next *Node
}

type LinkedList struct {
    Head *Node
    size int
}

func NewLinkedList() *LinkedList
func (l *LinkedList) Push(val int)          // add to front
func (l *LinkedList) Append(val int)        // add to back
func (l *LinkedList) Pop() (int, bool)      // remove from front
func (l *LinkedList) Len() int
func (l *LinkedList) ToSlice() []int        // front to back
func (l *LinkedList) Contains(val int) bool
func (l *LinkedList) Reverse()              // reverse in place
func (l *LinkedList) RemoveAt(index int) bool  // 0-indexed, returns false if out of range
```

## Examples

```
l := NewLinkedList()
l.Append(1)
l.Append(2)
l.Append(3)
l.ToSlice()  →  [1, 2, 3]

l.Push(0)
l.ToSlice()  →  [0, 1, 2, 3]

l.Pop()      →  0, true
l.ToSlice()  →  [1, 2, 3]

l.Reverse()
l.ToSlice()  →  [3, 2, 1]

l.RemoveAt(1)
l.ToSlice()  →  [3, 1]

l.Contains(3)  →  true
l.Contains(9)  →  false
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
go run ./cmd/drill check 176
```

or run the tests directly:

```bash
go test ./176-linked-list/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 176` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
