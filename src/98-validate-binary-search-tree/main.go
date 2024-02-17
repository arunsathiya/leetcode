package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isWithinRange(node, left, right *TreeNode) bool {
	if node == nil {
		return true
	}
	if left != nil && left.Val >= node.Val {
		return false
	}
	if right != nil && right.Val <= node.Val {
		return false
	}
	return isWithinRange(node.Left, left, node) && isWithinRange(node.Right, node, right)

}

func isValidBST(root *TreeNode) bool {
	return isWithinRange(root, nil, nil)
}
