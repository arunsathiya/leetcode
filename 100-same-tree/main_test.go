package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name   string
		first  *TreeNode
		second *TreeNode
		output bool
	}{
		{
			name:   "test case from leetcode",
			first:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			second: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := isSameTree(test.first, test.second)
			if output != test.output {
				t.Errorf("Expected %v, got %v", test.output, output)
			}
		})
	}
}
