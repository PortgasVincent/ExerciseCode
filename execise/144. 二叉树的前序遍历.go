package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traversal = func(node *TreeNode) {}
	traversal = func(node *TreeNode) {
		if node != nil {
			res = append(res, node.Val)
			traversal(node.Left)
			traversal(node.Right)
		}
	}
	traversal(root)
	return res
}
func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)
	return res
}

// 迭代
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}

// morris
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	for root != nil {
		p1 := root.Left
		if p1 != nil {
			for p1.Right != nil && p1.Right != root {
				p1 = p1.Right
			}
			if p1.Right == nil {
				res = append(res, root.Val)
				p1.Right = root
				root = root.Left
				continue
			}
			p1.Right = nil
			root = root.Right
		} else {
			res = append(res, root.Val)
			root = root.Right
		}

	}
	return res
}
