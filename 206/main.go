package main

import "fmt"

type ListNode struct {
     Val int
	Next *ListNode
 }

func reverseList(head *ListNode) *ListNode {
	var lastNode *ListNode
	var newHead *ListNode
	for {
		if head == nil {
			break
		}
		nextHeadTemp := head.Next
		newHead = head
		newHead.Next = lastNode
		lastNode = newHead
		head = nextHeadTemp
	}
	return newHead
}

func main() {
	node1 := &ListNode{1,nil}
	node2 := &ListNode{2,nil}
	node3 := &ListNode{3,nil}
	node1.Next = node2
	node2.Next = node3
	fmt.Println(reverseList(node1).Val)
}