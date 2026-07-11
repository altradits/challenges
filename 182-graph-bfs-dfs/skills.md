# Skills for 182-graph-bfs-dfs

## What You'll Learn

**Previous:** [181-binary-tree](../181-binary-tree/skills.md) | **Next:** [183-dynamic-programming](../183-dynamic-programming/skills.md)

**Challenge:** Represent a graph with an adjacency list. Traverse it with DFS and BFS.

## Graph Representation

An **adjacency list** maps each node to its neighbours:

```go
type Graph struct {
    Edges map[int][]int
}

// Undirected: edge (u, v) means both u→v and v→u
func (g *Graph) AddEdge(u, v int) {
    g.Edges[u] = append(g.Edges[u], v)
    g.Edges[v] = append(g.Edges[v], u)
}
```

Example:

```
Edges: {
    1: [2, 3],
    2: [1, 4],
    3: [1, 5],
    4: [2],
    5: [3],
}
```

## DFS (Depth-First Search)

Go as deep as possible before backtracking. Use recursion (implicit stack):

```go
func (g *Graph) DFS(start int) []int {
    visited := make(map[int]bool)
    order := []int{}

    var dfs func(node int)
    dfs = func(node int) {
        if visited[node] { return }
        visited[node] = true
        order = append(order, node)
        for _, n := range g.Edges[node] {
            dfs(n)
        }
    }

    dfs(start)
    return order
}
```

## BFS (Breadth-First Search)

Visit all nodes at distance 1 before distance 2. Use a queue:

```go
func (g *Graph) BFS(start int) []int {
    visited := map[int]bool{start: true}
    order := []int{}
    queue := []int{start}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        order = append(order, node)
        for _, n := range g.Edges[node] {
            if !visited[n] {
                visited[n] = true
                queue = append(queue, n)
            }
        }
    }
    return order
}
```

## When to Use DFS vs BFS

| Problem | Best Algorithm | Why |
|---------|---------------|-----|
| Find if path exists | DFS | Cheaper to implement |
| Shortest path (unweighted) | BFS | Explores nearest nodes first |
| Topological sort | DFS | Natural post-order |
| Level-by-level processing | BFS | Processes level by level |
| Connected components | DFS | Simple recursion |
| Maze solving | Either | BFS finds shortest |

## Documentation

- [Go Blog — Maps](https://go.dev/blog/maps)
- [Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) — weighted shortest path

**Next:** [183-dynamic-programming](../183-dynamic-programming/README.md)
