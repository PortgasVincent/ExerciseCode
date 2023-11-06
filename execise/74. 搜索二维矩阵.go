package main

import (
	"fmt"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	lx := len(matrix)
	ly := len(matrix[0])
	return search(matrix, target, 0, lx*ly)
}

func search(matrix [][]int, target, start, end int) bool {
	lx := len(matrix)
	ly := len(matrix[0])
	if start >= end-1 {
		x1, y1 := get(ly, start)
		x2, y2 := get(ly, end)
		if x1 >= 0 && x1 < lx && y1 >= 0 && y1 < ly && matrix[x1][y1] == target {
			return true
		}
		if x2 >= 0 && x2 < lx && y2 >= 0 && y2 < ly && matrix[x2][y2] == target {
			return true
		}
		return false
	}

	mid := (end + start) / 2
	x, y := get(ly, mid)
	if matrix[x][y] == target {
		return true
	}
	if matrix[x][y] < target {
		return search(matrix, target, mid+1, end)
	}
	return search(matrix, target, start, mid-1)
}

func get(l, sum int) (int, int) {
	x := sum / l
	y := sum%l - 1
	if y == -1 {
		x -= 1
		y = l - 1
	}
	return x, y
}
