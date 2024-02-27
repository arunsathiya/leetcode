package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{buildTree([]interface{}{1, 2, 3, 4, 5}), 3},
		{buildTree([]interface{}{1, 2}), 1},
	}

	for _, test := range tests {
		if got := diameterOfBinaryTree(test.root); got != test.expected {
			t.Errorf("diameterOfBinaryTree() = %d, want %d", got, test.expected)
		}
	}
}

func buildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	var root *TreeNode
	queue := []*TreeNode{}
	for i, val := range nodes {
		if val == nil {
			continue
		}
		newNode := &TreeNode{Val: val.(int)}
		if i == 0 {
			root = newNode
			queue = append(queue, newNode)
			continue
		}
		parent := queue[0]
		if parent.Left == nil {
			parent.Left = newNode
		} else if parent.Right == nil {
			parent.Right = newNode
			queue = append(queue[:0], queue[1:]...) // Dequeue
		}
		queue = append(queue, newNode)
	}
	return root
}
