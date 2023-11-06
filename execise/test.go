package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	temp := root
	for root != nil {
		p1 := root.Left
		if p1 != nil {
			for p1.Right != nil && p1.Right != root {
				p1 = p1.Right
			}
			if p1.Right == nil {
				p1.Right = root
				root = root.Left
				continue
			}
			p1.Right = nil
			res = addPath(res, root.Left)
		}
		root = root.Right
	}
	res = addPath(res, temp)
	return res
}
func addPath(res []int, root *TreeNode) []int {
	l1 := len(res)
	for root != nil {
		res = append(res, root.Val)
		root = root.Right
	}
	l2 := len(res) - 1
	for l1 < l2 {
		res[l1], res[l2] = res[l2], res[l1]
		l1++
		l2--
	}
	return res
}
