package main

import "testing"

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "Both lists with elements",
			list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			want:  &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			name:  "Both lists empty",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "Second list empty",
			list1: &ListNode{0, nil},
			list2: nil,
			want:  &ListNode{0, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !isEqual(got, tt.want) {
				t.Errorf("Test '%s' failed: linked list structure or values do not match expected result", tt.name)
			}
		})
	}
}

func isEqual(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
