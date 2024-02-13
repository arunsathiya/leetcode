package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name  string
		input *ListNode
		want  []int
	}{
		{
			name:  "5 elements",
			input: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			want:  []int{1, 5, 2, 4, 3},
		},
		{
			name:  "4 elements",
			input: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			want:  []int{1, 4, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReorderList(tt.input)
			if !reflect.DeepEqual(ListNodeToSlice(got), tt.want) {
				t.Errorf("Test %s failed: got %v, want %v", tt.name, ListNodeToSlice(got), tt.want)
			}
		})
	}
}

func ListNodeToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
