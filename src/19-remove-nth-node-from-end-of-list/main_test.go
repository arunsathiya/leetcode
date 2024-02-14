package main

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		n    int
		want *ListNode
	}{
		{
			name: "5 elements, remove 2nd from end",
			head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			n:    2,
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			name: "single element, remove 1st",
			head: &ListNode{1, nil},
			n:    1,
			want: nil,
		},
		{
			name: "2 elements, remove last",
			head: &ListNode{1, &ListNode{2, nil}},
			n:    1,
			want: &ListNode{1, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.head, tt.n)
			if !isEqual(got, tt.want) {
				t.Errorf("Test '%s' failed: linked list structure or values do not match expected result", tt.name)
			}
		})
	}
}

// isEqual compares two linked lists for equality in values and structure.
func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil // Both should be nil at the end for the lists to be equal
}
