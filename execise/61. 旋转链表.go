package main

import (
	"fmt"
)

func main() {
	fmt.Println(getPermutation(3, 1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	temp := head
	for temp != nil {
		l++
		temp = temp.Next
	}
	if k == 0 || l == 0 {
		return head
	}
	k = k % l

	count := 0
	var pre, cur *ListNode
	cur = head
	for count <= k {
		if count == k {
			pre = head
			break
		}
		cur = cur.Next
		count++
	}
	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}
	cur.Next = head
	head = pre.Next
	pre.Next = nil
	return head
}
