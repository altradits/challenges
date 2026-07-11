package piscine

type Queue struct {
	data []int
}

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *Queue) Front() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	return q.data[0], true
}

func (q *Queue) IsEmpty() bool { return len(q.data) == 0 }
func (q *Queue) Size() int     { return len(q.data) }

func BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	order := []int{}
	q := NewQueue()
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
