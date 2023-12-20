package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Length() int {
	if t == nil {
		return 0
	}
	return 1 + max(t.Left.Length(), t.Right.Length())
}

func (t *TreeNode) Swap() {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	t.Left.Swap()
	t.Right.Swap()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func invertTree(root *TreeNode) *TreeNode {
	if root.Length() == 0 {
		return &TreeNode{}
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
