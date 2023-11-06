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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	for left != nil {
		if left.Val >= root.Val {
			return false
		}
		left = left.Right
	}
	right := root.Right
	for right != nil {
		if right.Val <= root.Val {
			return false
		}
		right = right.Left
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
