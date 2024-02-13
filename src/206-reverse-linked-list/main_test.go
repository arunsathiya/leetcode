package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := map[string]struct {
		input *ListNode
		want  []int
	}{
		"reverse 5 elements": {&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, []int{5, 4, 3, 2, 1}},
		"reverse 2 elements": {&ListNode{1, &ListNode{2, nil}}, []int{2, 1}},
		"reverse empty list": {nil, []int{}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotList := ReverseList(tt.input)
			var got []int
			for gotList != nil {
				got = append(got, gotList.Val)
				gotList = gotList.Next
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}
}
