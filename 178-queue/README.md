# 178-queue — Queue (FIFO)

## Task

Implement a FIFO queue in `package piscine`:

```go
type Queue struct { /* unexported fields */ }

func NewQueue() *Queue
func (q *Queue) Enqueue(val int)
func (q *Queue) Dequeue() (int, bool)  // remove and return front element
func (q *Queue) Front() (int, bool)    // peek at front without removing
func (q *Queue) IsEmpty() bool
func (q *Queue) Size() int

// BFS returns nodes reachable from start in breadth-first order.
// graph is an adjacency list: graph[node] = list of neighbours.
func BFS(graph map[int][]int, start int) []int
```

## Examples

```
q := NewQueue()
q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)
q.Front()     →  1, true
q.Dequeue()   →  1, true
q.Dequeue()   →  2, true
q.Size()      →  1

graph := map[int][]int{
    1: {2, 3},
    2: {4},
    3: {4, 5},
    4: {},
    5: {},
}
BFS(graph, 1)  →  [1, 2, 3, 4, 5]
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
go run ./cmd/drill check 178
```

or run the tests directly:

```bash
go test ./178-queue/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 178` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
