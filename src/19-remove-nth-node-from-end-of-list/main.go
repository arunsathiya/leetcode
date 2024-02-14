package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var managed []*ListNode
	for head != nil {
		managed = append(managed, head)
		head = head.Next
	}
	indexToRemove := len(managed) - n
	if indexToRemove == 0 {
		return head.Next
	}
	if indexToRemove > 0 && indexToRemove < len(managed) {
		managed[indexToRemove-1].Next = managed[indexToRemove].Next
	} else if indexToRemove > 0 { // If removing the last node, ensure the previous node points to nil
		managed[indexToRemove-1].Next = nil
	}
	return head
}
