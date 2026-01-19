package main

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: true,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		got := isValidBST(tt.root)
		if got != tt.want {
			t.Errorf("isValidBST() = %v, want %v", got, tt.want)
		}
	}
}
