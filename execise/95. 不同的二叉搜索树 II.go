package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateTrees(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	if start == end {
		return []*TreeNode{
			{
				Val: start,
			},
		}
	}

	var res []*TreeNode
	for i := start; i <= end; i++ {
		leftSet := generate(start, i-1)
		rightSet := generate(i+1, end)
		for _, v1 := range leftSet {
			left := v1
			for _, v2 := range rightSet {
				root := &TreeNode{
					Val: i,
				}
				root.Left = left
				right := v2
				root.Right = right
				res = append(res, root)
			}
		}
	}
	return res
}
