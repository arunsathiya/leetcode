package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var node *ListNode
	for temp := head; temp != nil; temp = temp.Next {
		node = &ListNode{
			Val:  temp.Val,
			Next: node,
		}
	}
	return node
}
