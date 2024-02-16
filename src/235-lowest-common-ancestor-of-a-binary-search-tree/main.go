package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	current := root
	for current != nil {
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
		} else if p.Val > current.Val && q.Val > current.Val {
			current = current.Right
		} else {
			return current
		}
	}
	return current // but this is never reached
}
