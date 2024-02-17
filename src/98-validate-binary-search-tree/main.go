package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isWithinRange(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if (lower < node.Val) && (upper > node.Val) {
		return isWithinRange(node.Left, lower, node.Val) && isWithinRange(node.Right, node.Val, upper)
	} else {
		return false
	}
}

func isValidBST(root *TreeNode) bool {
	return isWithinRange(root, math.MinInt, math.MaxInt)
}
