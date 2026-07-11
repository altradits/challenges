# Skills for 178-queue

## What You'll Learn

**Previous:** [177-stack](../177-stack/skills.md) | **Next:** [179-binary-search](../179-binary-search/skills.md)

**Challenge:** Implement a FIFO queue and apply it to BFS graph traversal.

## What Is a Queue?

A queue is FIFO (First In, First Out) — like a line at a shop. Items are added at the back and removed from the front.

```
Enqueue(1) → [1]
Enqueue(2) → [1, 2]
Enqueue(3) → [1, 2, 3]
Dequeue()  → 1, [2, 3]
Dequeue()  → 2, [3]
```

## Slice-Backed Queue

```go
type Queue struct{ data []int }

func (q *Queue) Enqueue(val int) { q.data = append(q.data, val) }

func (q *Queue) Dequeue() (int, bool) {
    if len(q.data) == 0 { return 0, false }
    val := q.data[0]
    q.data = q.data[1:]  // removes front element; O(n) re-slice
    return val, true
}
```

**Note:** `q.data[1:]` creates a new slice header pointing to the second element but does not copy. For high-performance queues, use a circular buffer or `container/ring`.

## Breadth-First Search (BFS)

BFS visits all nodes at distance 1 before distance 2, using a queue to process them level by level:

```
Graph:
1 → [2, 3]
2 → [4]
3 → [4, 5]

BFS from 1:
Visit 1, enqueue [2, 3]
Visit 2, enqueue [4]
Visit 3, enqueue [4, 5] (4 already enqueued/visited — skip)
Visit 4
Visit 5
Order: [1, 2, 3, 4, 5]
```

```go
func BFS(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    order   := []int{}
    q       := NewQueue()

    q.Enqueue(start)
    visited[start] = true

    for !q.IsEmpty() {
        node, _ := q.Dequeue()
        order = append(order, node)
        for _, neighbour := range graph[node] {
            if !visited[neighbour] {
                visited[neighbour] = true
                q.Enqueue(neighbour)
            }
        }
    }
    return order
}
```

## Queue vs Stack in Graph Search

| Algorithm | Data Structure | Explores |
|-----------|---------------|---------|
| BFS | Queue (FIFO) | Nearest nodes first (shortest path) |
| DFS | Stack (LIFO) | Deep paths first |

## Documentation

- [container/ring](https://pkg.go.dev/container/ring) — circular buffer
- [container/list](https://pkg.go.dev/container/list) — efficient front/back operations

**Next:** [179-binary-search](../179-binary-search/README.md)
