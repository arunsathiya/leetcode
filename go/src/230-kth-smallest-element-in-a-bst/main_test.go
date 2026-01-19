package main

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			k:    1,
			want: 1,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			k:    3,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test with root %v and k %d", tt.root, tt.k), func(t *testing.T) {
			if got := kthSmallest(tt.root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
