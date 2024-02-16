package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root != nil && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	deque := list.New()
	deque.PushBack(root)
	var outer [][]int
	for deque.Len() > 0 {
		queueLength := deque.Len()
		var inner []int
		for i := 0; i < queueLength; i++ {
			node := deque.Front()
			if node != nil {
				inner = append(inner, int(node.Value))
				deque.PushBack(node.Prev())
				deque.PushBack(node.Next())
			}
		}
		if len(inner) > 0 {
			outer = append(outer, inner)
		}
	}
	return outer
}
