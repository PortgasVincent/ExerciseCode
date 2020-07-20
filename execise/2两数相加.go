package main

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	r := res
	var add bool
	temp := 0

	if l1 != nil {
		temp += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		temp += l2.Val
		l2 = l2.Next
	}
	if add {
		temp ++
	}
	if temp >= 10{
		add = true
		temp = temp % 10
	} else {
		add = false
	}
	r.Val = temp
	temp = 0
	for l1 != nil || l2 != nil || add == true {
		r.Next = &ListNode{}
		r = r.Next
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		if add {
			temp ++
		}
		if temp >= 10{
			add = true
			temp = temp % 10
		} else {
			add = false
		}
		r.Val = temp
		temp = 0
	}
	return res
}

func mainfalse(){}