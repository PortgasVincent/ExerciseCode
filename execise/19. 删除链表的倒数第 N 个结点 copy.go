package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre, now *ListNode
	now = head
	var count int
	for now.Next != nil {
		count++
		now = now.Next
		if count == n {
			pre = head
			break
		}
	}
	for now.Next != nil {
		now = now.Next
		pre = pre.Next
	}
	if pre != nil && pre.Next != nil {
		pre.Next = pre.Next.Next
	}

	if pre == nil {
		return head.Next
	}
	return head
}
