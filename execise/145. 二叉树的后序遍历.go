package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = postorderTraversal(root.Left)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 迭代
func postorderTraversal(root *TreeNode) []int {
	var res = []int{}
	var stack = []*TreeNode{}
	var pre *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || pre == root.Right {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}

// Morris
func postorderTraversal(root *TreeNode) []int {
	var res = []int{}
	addPath := func(node *TreeNode) {
		l := len(res)
		for node != nil {
			res = append(res, node.Val)
			node = node.Right
		}
		reserve(res[l:])
	}
	p1 := root
	for p1 != nil {
		p2 := p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return res
}

func reserve(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return
}
