package main

import (
	"fmt"
)

func main() {
	lists := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	fmt.Println(deleteDuplicates(lists))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	next := cur.Next
	for next != nil {
		if next.Val != cur.Val {
			cur.Next = next
			cur = cur.Next
		}
		next = next.Next
	}
	cur.Next = nil
	return head
}
