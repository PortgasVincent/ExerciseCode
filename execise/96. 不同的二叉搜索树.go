package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numTrees(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	m := map[string]int{}
	return num(1, n, m)
}

func num(start, end int, m map[string]int) int {
	if start >= end {
		return 1
	}
	if v, ok := m[getKey(start, end)]; ok {
		return v
	}

	var res int = 0
	for i := start; i <= end; i++ {
		left := num(start, i-1, m)
		right := num(i+1, end, m)
		res = res + left*right
	}
	m[getKey(start, end)] = res
	return res
}

func getKey(a, b int) string {
	return strconv.Itoa(a) + "," + strconv.Itoa(b)
}
