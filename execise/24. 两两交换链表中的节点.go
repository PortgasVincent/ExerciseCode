package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var cur, prev, next, temp *ListNode
	cur = head
	if cur == nil || cur.Next == nil {
		return head
	}
	next = cur.Next
	temp = next.Next
	head = next
	head.Next = cur
	cur.Next = temp
	// head(next) -> cur -> temp
	// 						-> prev -> cur -> next
	for cur.Next != nil && cur.Next.Next != nil {
		prev = cur
		cur = cur.Next
		next = cur.Next

		temp = next.Next
		prev.Next = next
		next.Next = cur
		cur.Next = temp
	}
	return head
}
