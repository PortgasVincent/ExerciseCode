package main

import (
	"fmt"
)

func main() {
	lists := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	res := partition(lists, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var fir, sec, p, q *ListNode
	for head != nil {
		if head.Val < x {
			if fir == nil {
				fir = head
				p = head
			} else {
				p.Next = head
				p = p.Next
			}
		} else {
			if sec == nil {
				sec = head
				q = head
			} else {
				q.Next = head
				q = q.Next
			}
		}

		head = head.Next
	}
	if p != nil {
		p.Next = sec
	}
	if q != nil {
		q.Next = nil
	}
	if fir == nil {
		return sec
	}
	return fir
}
