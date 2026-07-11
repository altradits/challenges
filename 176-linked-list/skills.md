# Skills for 176-linked-list

## What You'll Learn

**Previous:** [169-testing-advanced](../169-testing-advanced/skills.md) | **Next:** [177-stack](../177-stack/skills.md)

**Challenge:** Build a linked list from scratch — the most fundamental dynamic data structure.

## What Is a Linked List?

A linked list is a sequence of **nodes**, each holding a value and a pointer to the next node.

```
Head
  ↓
[1] → [2] → [3] → nil
```

Unlike a slice, nodes are **not contiguous in memory**. Each node holds the address of the next. This makes prepend O(1) (just update Head) but random access O(n) (must walk from Head).

## Node and List Structure

```go
type Node struct {
    Val  int
    Next *Node
}

type LinkedList struct {
    Head *Node
    size int
}
```

`Head` points to the first node. `nil` means an empty list.

## Core Operations

### Push (prepend) — O(1)

```go
func (l *LinkedList) Push(val int) {
    l.Head = &Node{Val: val, Next: l.Head}
    l.size++
}
```

Before: `Head → [1] → [2] → nil`
After `Push(0)`: `Head → [0] → [1] → [2] → nil`

### Append — O(n)

Walk to the last node, then attach:

```go
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
```

### Pop (remove from front) — O(1)

```go
func (l *LinkedList) Pop() (int, bool) {
    if l.Head == nil {
        return 0, false
    }
    val := l.Head.Val
    l.Head = l.Head.Next
    l.size--
    return val, true
}
```

### Reverse — O(n)

Walk the list, reversing each pointer:

```go
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
```

Visual:

```
Before: nil ← [1] → [2] → [3] → nil
Step 1: nil ← [1]   [2] → [3] → nil  (curr.Next = prev)
Step 2: nil ← [1] ← [2]   [3] → nil
Step 3: nil ← [1] ← [2] ← [3]   nil
Head = prev = [3]
```

## Time Complexity

| Operation | Time | Why |
|-----------|------|-----|
| Push | O(1) | Update Head only |
| Pop | O(1) | Update Head only |
| Append | O(n) | Walk to last node |
| Contains | O(n) | Walk until found |
| RemoveAt(i) | O(n) | Walk to index i-1 |
| Reverse | O(n) | Walk once |
| Len | O(1) | Maintained in `size` |

## When to Use a Linked List

| Prefer linked list | Prefer slice |
|-------------------|--------------|
| Frequent prepend/pop from front | Random access by index |
| Unknown size, many inserts | Iteration over all elements |
| Implementing a stack or queue | Cache-friendly sequential reads |

In practice, Go slices (backed by arrays) often win due to cache locality — measure before choosing.

## Solving the Challenge

```go
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
    for curr.Next != nil { curr = curr.Next }
    curr.Next = newNode
    l.size++
}

func (l *LinkedList) Pop() (int, bool) {
    if l.Head == nil { return 0, false }
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
        if curr.Val == val { return true }
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
    if index < 0 || index >= l.size { return false }
    if index == 0 {
        l.Head = l.Head.Next
        l.size--
        return true
    }
    curr := l.Head
    for i := 0; i < index-1; i++ { curr = curr.Next }
    curr.Next = curr.Next.Next
    l.size--
    return true
}
```

## Documentation

- [Go Spec — Pointer types](https://go.dev/ref/spec#Pointer_types)
- [container/list — Go standard library doubly linked list](https://pkg.go.dev/container/list)
- [Data Structures and Algorithms in Go (book)](https://www.packtpub.com/product/data-structures-and-algorithms-with-go/9781789618501)

**Next:** [177-stack](../177-stack/README.md)
