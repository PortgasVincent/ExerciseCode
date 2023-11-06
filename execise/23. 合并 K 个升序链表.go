package main

import "fmt"

func main() {
	list := []*ListNode{
		nil,
		nil,
	}
	res := mergeKLists(list)
	fmt.Println("res: ")
	printL([]*ListNode{res})

}

func printL(list []*ListNode) {
	for _, v := range list {
		var node *ListNode = v
		for node != nil {
			fmt.Println(node.Val)
			node = node.Next
		}
		fmt.Println("---")
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	k := l / 2
	return merge(lists[:k], lists[k:])[0]
}

func merge(list1, list2 []*ListNode) []*ListNode {
	l1 := len(list1)
	l2 := len(list2)
	k1 := l1 / 2
	k2 := l2 / 2

	if l1 == 0 || (l1 == 1 && list1[0] == nil) {
		if l2 <= 1 {
			return list2
		}
		return merge(list2[:k2], list2[k2:])
	}
	if l2 == 0 || (l2 == 1 && list2[0] == nil) {
		if l1 <= 1 {
			return list1
		}
		return merge(list1[:k1], list1[k1:])
	}
	if l1 == 1 && l2 == 1 {
		var newList *ListNode
		if list1[0].Val <= list2[0].Val {
			newList = list1[0]
			list1[0] = list1[0].Next
		} else {
			newList = list2[0]
			list2[0] = list2[0].Next
		}
		temp := newList
		for list1[0] != nil && list2[0] != nil {
			if list1[0].Val <= list2[0].Val {
				temp.Next = list1[0]
				list1[0] = list1[0].Next
				temp = temp.Next
			} else {
				temp.Next = list2[0]
				list2[0] = list2[0].Next
				temp = temp.Next
			}
		}
		if list1[0] != nil {
			temp.Next = list1[0]
		}
		if list2[0] != nil {
			temp.Next = list2[0]
		}
		return []*ListNode{newList}
	}

	if l1 > 1 {
		list1 = merge(list1[:k1], list1[k1:])
	}
	if l2 > 1 {
		list2 = merge(list2[:k2], list2[k2:])
	}
	return merge(list1, list2)
}
