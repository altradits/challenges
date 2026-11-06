package piscine

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func NewLinkedList() *LinkedList { return &LinkedList{} }

func (l *LinkedList) Push(val int) {
	l.Head = &Node{Val: val, Next: l.Head}
	l.size++
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{Val: val}
	if l.Head == nil {
		l.Head = newNode
		l.size++
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	l.size++
}

func (l *LinkedList) Pop() (int, bool) {
	if l.Head == nil {
		return 0, false
	}
	val := l.Head.Val
	l.Head = l.Head.Next
	l.size--
	return val, true
}

func (l *LinkedList) Len() int { return l.size }

func (l *LinkedList) ToSlice() []int {
	result := make([]int, 0, l.size)
	for curr := l.Head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}

func (l *LinkedList) Contains(val int) bool {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return true
		}
	}
	return false
}

func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func (l *LinkedList) RemoveAt(index int) bool {
	if index < 0 || index >= l.size {
		return false
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.size--
		return true
	}
	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	l.size--
	return true
}
