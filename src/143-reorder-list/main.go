package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var managed []*ListNode
	for head != nil {
		managed = append(managed, head)
		head = head.Next
	}
	i, j := 0, len(managed)-1
	for i < j {
		managed[i].Next = managed[j]
		i++
		if i == j {
			break
		}
		managed[j].Next = managed[i]
		j--
	}
	managed[i].Next = nil
}
