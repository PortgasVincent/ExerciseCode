package main

import (
	"fmt"
)

func main() {
	list := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isValidBST(list))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var r1, r2, pre *TreeNode
	for root != nil {
		if root.Left != nil {
			p := root.Left
			for p.Right != nil && p.Right != root {
				p = p.Right
			}
			if p.Right == nil {
				p.Right = root
				root = root.Left
			} else {
				if pre != nil && pre.Val > root.Val {
					r1 = root
					if r2 == nil {
						r2 = pre
					}
				}
				pre = root
				p.Right = nil
				root = root.Right
			}
		} else {
			if pre != nil && pre.Val > root.Val {
				r1 = root
				if r2 == nil {
					r2 = pre
				}
			}
			pre = root
			root = root.Right
		}
	}
	r1.Val, r2.Val = r2.Val, r1.Val
	return
}
