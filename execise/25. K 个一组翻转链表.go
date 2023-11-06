package main

import "fmt"

func main() {
	h, _ := reverse(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	fir := head
	temp := head
	count := 0
	var preTemp *ListNode
	var h, t, prev *ListNode
	for temp != nil && count < k {
		preTemp = temp
		temp = temp.Next
		count++
	}
	if count < k {
		return head
	}
	preTemp.Next = nil
	h, t = reverse(fir)
	head = h
	t.Next = temp
	prev = t

	for temp != nil {
		fir = temp
		count = 0
		for temp != nil && count < k {
			preTemp = temp
			temp = temp.Next
			count++
		}
		if count < k {
			return head
		}
		preTemp.Next = nil
		h, t = reverse(fir)
		prev.Next = h
		t.Next = temp
		prev = t
	}
	return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev, head
}
