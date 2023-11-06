package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var p, head1, pre, next *ListNode
	i := 1
	cur := head
	for i < left {
		if i == left-1 {
			p = cur
		}
		cur = cur.Next
		i++
	}
	head1 = cur

	for i < right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		i++
	}
	next = cur.Next
	cur.Next = pre
	head1.Next = next
	if left == 1 {
		return cur
	}

	p.Next = cur
	return head
}
