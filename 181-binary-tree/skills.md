# Skills for 181-binary-tree

## What You'll Learn

**Previous:** [180-sorting-algorithms](../180-sorting-algorithms/skills.md) | **Next:** [182-graph-bfs-dfs](../182-graph-bfs-dfs/skills.md)

**Challenge:** Build a binary search tree with insertion, search, and three traversals using recursion.

## Binary Search Tree Property

A BST maintains one invariant at every node:
- All values in the **left subtree** are **less than** the node's value
- All values in the **right subtree** are **greater than** the node's value

```
        5
       / \
      3   7
     / \
    1   4
```

This invariant makes search O(log n) on a balanced tree.

## Recursive Insert

```go
func insertNode(node *TreeNode, val int) *TreeNode {
    if node == nil {
        return &TreeNode{Val: val}  // base case: create leaf
    }
    if val < node.Val {
        node.Left = insertNode(node.Left, val)
    } else if val > node.Val {
        node.Right = insertNode(node.Right, val)
    }
    return node  // return unchanged node (or newly created one)
}
```

This pattern — "return the (possibly new) node" — avoids needing a parent pointer.

## Three Traversals

The three traversal orders differ only in when you visit the current node:

```go
// In-Order: Left, Root, Right → produces sorted output for BST
func inOrder(node *TreeNode, result *[]int) {
    if node == nil { return }
    inOrder(node.Left, result)
    *result = append(*result, node.Val)  // visit here
    inOrder(node.Right, result)
}

// Pre-Order: Root, Left, Right → used to copy/serialise tree
func preOrder(node *TreeNode, result *[]int) {
    if node == nil { return }
    *result = append(*result, node.Val)  // visit here
    preOrder(node.Left, result)
    preOrder(node.Right, result)
}

// Post-Order: Left, Right, Root → used to delete tree
func postOrder(node *TreeNode, result *[]int) {
    if node == nil { return }
    postOrder(node.Left, result)
    postOrder(node.Right, result)
    *result = append(*result, node.Val)  // visit here
}
```

Using a **closure** instead of pointer to slice is cleaner (see solution).

## Closure Pattern for Tree Traversal

```go
func (t *BST) InOrder() []int {
    result := []int{}
    var traverse func(*TreeNode)
    traverse = func(node *TreeNode) {
        if node == nil { return }
        traverse(node.Left)
        result = append(result, node.Val)
        traverse(node.Right)
    }
    traverse(t.Root)
    return result
}
```

The inner `traverse` captures `result` by reference — no need to pass `&result`.

## Height (Max Depth)

```go
func height(node *TreeNode) int {
    if node == nil { return 0 }
    left  := height(node.Left)
    right := height(node.Right)
    if left > right { return left + 1 }
    return right + 1
}
```

## Time Complexity

| Operation | Balanced BST | Worst Case (sorted input) |
|-----------|-------------|--------------------------|
| Insert | O(log n) | O(n) — degenerates to linked list |
| Contains | O(log n) | O(n) |
| InOrder | O(n) | O(n) |

For guaranteed O(log n), use a self-balancing tree (AVL, Red-Black). Go's standard library provides none — use a sorted slice with `sort.SearchInts` or the `btree` package for B-trees.

## Documentation

- [Go Spec — Recursive data types](https://go.dev/ref/spec#Types)
- [container/heap](https://pkg.go.dev/container/heap) — heap (priority queue) in Go stdlib
- [google/btree](https://github.com/google/btree) — B-tree implementation

**Next:** [182-graph-bfs-dfs](../182-graph-bfs-dfs/README.md)
