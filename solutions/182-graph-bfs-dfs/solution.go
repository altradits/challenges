package piscine

type Graph struct {
	Edges map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Edges: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)
	g.Edges[v] = append(g.Edges[v], u)
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	order := []int{}
	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		order = append(order, node)
		for _, neighbour := range g.Edges[node] {
			dfs(neighbour)
		}
	}
	dfs(start)
	return order
}

func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	order := []int{}
	queue := []int{start}
	visited[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, neighbour := range g.Edges[node] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	return order
}

func (g *Graph) HasPath(src, dst int) bool {
	visited := make(map[int]bool)
	var dfs func(node int) bool
	dfs = func(node int) bool {
		if node == dst {
			return true
		}
		if visited[node] {
			return false
		}
		visited[node] = true
		for _, neighbour := range g.Edges[node] {
			if dfs(neighbour) {
				return true
			}
		}
		return false
	}
	return dfs(src)
}

func (g *Graph) ConnectedComponents() int {
	visited := make(map[int]bool)
	components := 0
	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, neighbour := range g.Edges[node] {
			dfs(neighbour)
		}
	}
	for node := range g.Edges {
		if !visited[node] {
			dfs(node)
			components++
		}
	}
	return components
}
