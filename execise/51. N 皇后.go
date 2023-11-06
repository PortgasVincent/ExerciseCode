package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

var solution = [][]string{}

func solveNQueens(n int) [][]string {
	solution = [][]string{} // 避免多个测试用例之间影响
	queen := make([]int, n)
	for i := 0; i < n; i++ {
		queen[i] = -1
	}
	col := map[int]bool{}
	xie1 := map[int]bool{}
	xie2 := map[int]bool{}

	back(0, n, queen, col, xie1, xie2)
	return solution
}

func back(row, n int, queen []int, col, xie1, xie2 map[int]bool) {
	if row == n {
		sub := record(n, queen)
		solution = append(solution, sub) // 全局变量无需返回值
		return
	}
	for i := 0; i < n; i++ {
		if col[i] || xie1[row-i] || xie2[row+i] {
			continue
		}
		queen[row] = i
		col[i] = true
		xie1[row-i] = true
		xie2[row+i] = true
		back(row+1, n, queen, col, xie1, xie2)

		queen[row] = -1
		col[i] = false
		xie1[row-i] = false
		xie2[row+i] = false
	}
	return
}

func record(n int, queen []int) []string {
	var sub = []string{}
	for _, v := range queen {
		str := ""
		for i := 0; i < n; i++ {
			if i == v {
				str += "Q"
				continue
			}
			str += "."
		}
		sub = append(sub, str)
	}
	return sub
}
