package main

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "5 elements, remove 2nd from end",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "single element, remove 1st",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "2 elements, remove last",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToListNode(tt.head)
			got := removeNthFromEnd(head, tt.n)
			gotSlice := listNodeToSlice(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("Test '%s' failed: got %v, want %v", tt.name, gotSlice, tt.want)
			}
		})
	}
}

func sliceToListNode(slice []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, val := range slice {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

func listNodeToSlice(head *ListNode) []int {
	var slice []int
	for current := head; current != nil; current = current.Next {
		slice = append(slice, current.Val)
	}
	return slice
}
