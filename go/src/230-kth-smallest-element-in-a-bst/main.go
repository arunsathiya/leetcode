package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(root *TreeNode, c chan int) {
	if root == nil {
		return
	}
	walk(root.Left, c)
	c <- root.Val
	walk(root.Right, c)
}

func kthSmallest(root *TreeNode, k int) int {
	c := make(chan int)
	go walk(root, c)
	for i := 0; i < k-1; i++ {
		<-c
	}
	return <-c
}
