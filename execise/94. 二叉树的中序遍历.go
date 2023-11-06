package main

import (
	"fmt"
)

func main() {
	list := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(list))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // 递归
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = inorderTraversal(root.Left)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

// 迭代
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// Morris
func inorderTraversal(root *TreeNode) []int {
	var res = []int{}
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
				res = append(res, root.Val)
				p.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
