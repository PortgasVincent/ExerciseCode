package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, has := check(root)
	return !has
}

//[3,9,20,null,null,15,7]
func check(node *TreeNode)(int, bool){
	if node == nil{
		return 0, false
	}
	left, leftHas := check(node.Left)
	right, rightHas := check(node.Right)
	if leftHas || rightHas {
		return 0, true
	}
	m, l := tenary(left, right)
	if l > 1{
		return 0, true
	}
	return m+1, false
}

func tenary(a, b int)(int, int){
	if a > b{
		return a, a-b
	}
	return b, b-a
}