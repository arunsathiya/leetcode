package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var managed []*ListNode
	for head != nil {
		managed = append(managed, head)
		head = head.Next
	}
	i, j := 0, len(managed)-1
	for i < j {

	}
}
