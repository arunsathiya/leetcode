package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		inOrder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return result
}
