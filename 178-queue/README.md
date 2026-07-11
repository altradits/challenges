# 178-queue — Queue (FIFO)

## Challenge

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
