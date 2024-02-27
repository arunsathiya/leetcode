package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	count := 0
	originalRoot := root
	for root.Left != nil {
		count++
		root = root.Left
	}
	root = originalRoot
	for root.Right != nil {
		count++
		root = root.Right
	}
	return count
}
