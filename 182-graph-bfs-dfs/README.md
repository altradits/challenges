# 182-graph-bfs-dfs — Graph Traversal

## Challenge

Implement in `package piscine`:

```go
// Graph is an undirected graph represented as an adjacency list.
type Graph struct {
    Edges map[int][]int
}

func NewGraph() *Graph
func (g *Graph) AddEdge(u, v int)

// DFS returns nodes reachable from start in depth-first order.
func (g *Graph) DFS(start int) []int

// BFS returns nodes reachable from start in breadth-first order.
func (g *Graph) BFS(start int) []int

// HasPath reports whether a path exists from src to dst.
func (g *Graph) HasPath(src, dst int) bool

// ConnectedComponents returns the number of connected components.
func (g *Graph) ConnectedComponents() int
```

## Examples

```
g := NewGraph()
g.AddEdge(1, 2)
g.AddEdge(1, 3)
g.AddEdge(2, 4)
g.AddEdge(3, 5)

g.DFS(1)  →  [1, 2, 4, 3, 5]   (or any valid DFS order)
g.BFS(1)  →  [1, 2, 3, 4, 5]   (or any valid BFS order)

g.HasPath(1, 5)  →  true
g.HasPath(4, 5)  →  false  (4 and 5 not directly connected)

// Disconnected graph
g2 := NewGraph()
g2.AddEdge(1,2)
g2.AddEdge(3,4)
g2.ConnectedComponents()  →  2
```

Read `skills.md` before you start.
