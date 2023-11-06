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

	var now int = -101
	for head.Next != nil && (head.Val == head.Next.Val || head.Val == now) {
		now = head.Val
		head = head.Next
	}
	if head.Next == nil {
		if head.Val == now {
			return nil
		}
		return head
	}

	cur := head
	p := head.Next
	q := p.Next
	for q != nil {
		if p.Val == q.Val {
			now = p.Val
			p = q
			q = q.Next
			continue
		}
		if p.Val != now {
			cur.Next = p
			cur = cur.Next
		}
		p = q
		q = q.Next
	}
	if p.Val != now {
		cur.Next = p
		cur = cur.Next
	}
	cur.Next = nil
	return head
}
