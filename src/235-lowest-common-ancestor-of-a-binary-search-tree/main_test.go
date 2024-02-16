package main

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		pVal     int // Use value for comparison
		qVal     int // Use value for comparison
		expected int // Use value for expected result
	}{
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			pVal:     2,
			qVal:     8,
			expected: 6,
		},
		{
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			pVal:     2,
			qVal:     4,
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			pVal:     2,
			qVal:     1,
			expected: 2,
		},
	}

	for _, test := range tests {
		p := &TreeNode{Val: test.pVal}
		q := &TreeNode{Val: test.qVal}
		actual := lowestCommonAncestor(test.root, p, q)
		if actual == nil || actual.Val != test.expected {
			t.Errorf("Test failed: for p = %d, q = %d, expected %d, got %v", test.pVal, test.qVal, test.expected, actual)
		}
	}
}
