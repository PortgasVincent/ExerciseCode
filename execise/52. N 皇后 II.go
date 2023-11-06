package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

var solution [][]string

func totalNQueens(n int) int {
	solution = [][]string{}
	queen := make([]int, n)
	for i := 0; i < n; i++ {
		queen[i] = -1
	}
	col := map[int]bool{}
	xie1 := map[int]bool{}
	xie2 := map[int]bool{}

	back(queen, 0, n, col, xie1, xie2)
	return len(solution)
}

func back(queen []int, row, n int, col, xie1, xie2 map[int]bool) {
	if row == n {
		record(queen, n)
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
		back(queen, row+1, n, col, xie1, xie2)

		queen[row] = -1
		col[i] = false
		xie1[row-i] = false
		xie2[row+i] = false
	}
}

func record(queen []int, n int) {
	sub := []string{}
	for _, v := range queen {
		var str = ""
		for i := 0; i < n; i++ {
			if i == v {
				str += "Q"
				continue
			}
			str += "."
		}
		sub = append(sub, str)
	}
	solution = append(solution, sub)
	return
}
