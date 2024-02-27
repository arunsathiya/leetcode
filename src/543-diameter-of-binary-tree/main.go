package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		max = maxHelper(max, leftDepth+rightDepth)
		return 1 + maxHelper(leftDepth, rightDepth)
	}
	depth(root)
	return max
}

func maxHelper(a, b int) int {
	if a > b {
		return a
	}
	return b
}
