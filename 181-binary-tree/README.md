# 181-binary-tree — Binary Search Tree

## Task

Implement a Binary Search Tree (BST) in `package piscine`:

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type BST struct {
    Root *TreeNode
}

func NewBST() *BST
func (t *BST) Insert(val int)
func (t *BST) Contains(val int) bool
func (t *BST) InOrder() []int      // returns values sorted ascending
func (t *BST) PreOrder() []int     // root, left, right
func (t *BST) PostOrder() []int    // left, right, root
func (t *BST) Height() int         // max depth from root to any leaf
func (t *BST) Size() int           // total number of nodes
```

## Examples

```
t := NewBST()
t.Insert(5)
t.Insert(3)
t.Insert(7)
t.Insert(1)
t.Insert(4)

t.Contains(3)    →  true
t.Contains(9)    →  false
t.InOrder()      →  [1, 3, 4, 5, 7]
t.PreOrder()     →  [5, 3, 1, 4, 7]
t.PostOrder()    →  [1, 4, 3, 7, 5]
t.Height()       →  3
t.Size()         →  5
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
go run ./cmd/drill check 181
```

or run the tests directly:

```bash
go test ./181-binary-tree/
```

Green means done. Red means read the failure message, fix from memory, check
again.

## Repeat schedule

Redo this challenge from scratch after **1 day, 3 days, 7 days, 21 days**.
Each round, start the timer with `go run ./cmd/drill start 181` and beat your
previous time. `go run ./cmd/drill due` tells you what is due today.
