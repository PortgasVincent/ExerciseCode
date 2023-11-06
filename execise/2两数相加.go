package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	res := &ListNode{}
// 	r := res
// 	var add bool
// 	temp := 0

// 	if l1 != nil {
// 		temp += l1.Val
// 		l1 = l1.Next
// 	}
// 	if l2 != nil {
// 		temp += l2.Val
// 		l2 = l2.Next
// 	}
// 	if add {
// 		temp ++
// 	}
// 	if temp >= 10{
// 		add = true
// 		temp = temp % 10
// 	} else {
// 		add = false
// 	}
// 	r.Val = temp
// 	temp = 0
// 	for l1 != nil || l2 != nil || add == true {
// 		r.Next = &ListNode{}
// 		r = r.Next
// 		if l1 != nil {
// 			temp += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil {
// 			temp += l2.Val
// 			l2 = l2.Next
// 		}
// 		if add {
// 			temp ++
// 		}
// 		if temp >= 10{
// 			add = true
// 			temp = temp % 10
// 		} else {
// 			add = false
// 		}
// 		r.Val = temp
// 		temp = 0
// 	}
// 	return res
// }

func main() {
	l1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	temp := res
	var largeTen bool

	var add1, add2 int
	if l1 != nil {
		add1 = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		add2 = l2.Val
		l2 = l2.Next
	}
	temp.Val = add1 + add2
	if temp.Val >= 10 {
		temp.Val = temp.Val % 10
		largeTen = true
	}
	for l1 != nil || l2 != nil {
		temp.Next = &ListNode{}
		temp = temp.Next

		var add3, add4 int
		if l1 != nil {
			add3 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			add4 = l2.Val
			l2 = l2.Next
		}
		temp.Val = add3 + add4
		if largeTen {
			temp.Val++
			largeTen = false
		}
		if temp.Val >= 10 {
			temp.Val = temp.Val % 10
			largeTen = true
		}
	}
	if largeTen {
		temp.Next = &ListNode{Val: 1}
	}
	return res
}
