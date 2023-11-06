package main

import (
	"fmt"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func setZeroes(matrix [][]int) {
	lx := len(matrix)
	ly := len(matrix[0])
	var col0 bool
	for i := 0; i < lx; i++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1; j < ly; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := lx - 1; i >= 0; i-- {
		for j := 1; j < ly; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if col0 {
		for i := 0; i < lx; i++ {
			matrix[i][0] = 0
		}
	}
	return
}
