package main

import (
	"testing"
)

func isEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func TestReverseList(t *testing.T) {
	tests := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"reverse 5 elements": {&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}},
		"reverse 2 elements": {&ListNode{1, &ListNode{2, nil}}, &ListNode{2, &ListNode{1, nil}}},
		"reverse empty list": {nil, nil},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := ReverseList(tt.input)
			if !isEqual(got, tt.want) {
				t.Errorf("%s: lists are not equal", name)
			}
		})
	}
}
