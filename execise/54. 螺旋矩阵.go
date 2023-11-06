package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	xL, yL := len(matrix), len(matrix[0])
	if xL == 1 {
		return matrix[0]
	}
	y1, x2, y2, x1 := yL-1, xL-1, 0, 0
	var res []int
	x, y := 0, 0
	for x1 < x2 && y2 < y1 {
		for ; y < y1; y++ {
			res = append(res, matrix[x][y])
		}
		y1--
		for ; x < x2; x++ {
			res = append(res, matrix[x][y])
		}
		x2--
		for ; y > y2; y-- {
			res = append(res, matrix[x][y])
		}
		y2++
		for ; x > x1; x-- {
			res = append(res, matrix[x][y])
		}
		x1++
		x++
		y++
	}
	if x1 == x2 && y1 == y2 {
		res = append(res, matrix[x1][y1])
	} else if x1 == x2 {
		for y := y2; y <= y1; y++ {
			res = append(res, matrix[x1][y])
		}
	} else if y1 == y2 {
		for x := x1; x <= x2; x++ {
			res = append(res, matrix[x][y1])
		}
	}
	return res
}
