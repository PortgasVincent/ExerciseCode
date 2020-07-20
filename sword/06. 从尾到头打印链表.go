package main

func main(){}

type ListNode struct {
	Val int
    Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	l := len(res)-1
	for i := 0; i < l - i;i++ {
		temp := res[i]
		res[i] = res[l-i]
		res[l-i] = temp
	}
	return res
}