package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			root: nil,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test with root %v", tt.root), func(t *testing.T) {
			got := levelOrder(tt.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
