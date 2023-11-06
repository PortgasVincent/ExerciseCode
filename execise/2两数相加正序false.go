package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 错误答案❌
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	a := []int{}
	b := []int{}
	l3 := l1
	l4 := l2
	for i := 0; l3 != nil; i++ {
		a = append(a, l3.Val)
		l3 = l3.Next
	}
	for i := 0; l4 != nil; i++ {
		b = append(b, l4.Val)
		l4 = l4.Next
	}
	lenA := len(a)
	lenB := len(b)
	res := &ListNode{}
	r := res
	if lenA < lenB {
		c := a
		a = b
		b = c

		lenC := lenA
		lenA = lenB
		lenB = lenC
	}

	diff := lenA - lenB
	var add bool
	var temp int

	if lenA == lenB {
		temp = a[0] + b[0]

	} else {
		temp = a[0]
	}
	if temp >= 10 {
		add = true
		temp = temp % 10
	}
	r.Val = temp

	for i := 1; i < lenA || add == true; i++ {
		r.Next = &ListNode{}
		r = r.Next

		// reach max limit and add == true
		if i == lenA {
			r.Val = 1
			break
		}
		if i >= diff {
			temp = a[i] + b[i-diff]
		} else {
			temp = a[i]
		}
		if add {
			temp++
		}
		if temp >= 10 {
			add = true
			temp = temp % 10
		} else {
			add = false
		}
		r.Val = temp

	}
	return res
}

func main2() {

}
