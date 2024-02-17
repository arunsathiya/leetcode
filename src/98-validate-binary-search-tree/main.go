package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != nil && prev.Val >= root.Val {
				return false
			}
			prev = root
			root = root.Right
		}
	}
	return true
}
