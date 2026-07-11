# 181-binary-tree — Binary Search Tree

## Challenge

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
