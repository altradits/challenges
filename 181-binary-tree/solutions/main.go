package piscine

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func NewBST() *BST { return &BST{} }

func (t *BST) Insert(val int) {
	t.Root = insertNode(t.Root, val)
}

func insertNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		node.Left = insertNode(node.Left, val)
	} else if val > node.Val {
		node.Right = insertNode(node.Right, val)
	}
	return node
}

func (t *BST) Contains(val int) bool {
	curr := t.Root
	for curr != nil {
		if val == curr.Val {
			return true
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return false
}

func (t *BST) InOrder() []int {
	result := []int{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(t.Root)
	return result
}

func (t *BST) PreOrder() []int {
	result := []int{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(t.Root)
	return result
}

func (t *BST) PostOrder() []int {
	result := []int{}
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}
	traverse(t.Root)
	return result
}

func (t *BST) Height() int {
	return height(t.Root)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	right := height(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func (t *BST) Size() int {
	return size(t.Root)
}

func size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.Left) + size(node.Right)
}
