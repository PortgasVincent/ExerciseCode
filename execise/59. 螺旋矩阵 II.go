package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	var y2, x2, y1, x1 = n - 1, n - 1, 0, 0
	x, y, count := 0, 0, 0
	for x1 < x2 {
		for ; y < y2; y++ {
			count++
			res[x][y] = count
		}
		y2--
		for ; x < x2; x++ {
			count++
			res[x][y] = count
		}
		x2--
		for ; y > y1; y-- {
			count++
			res[x][y] = count
		}
		y1++
		for ; x > x1; x-- {
			count++
			res[x][y] = count
		}
		x1++
		x++
		y++
	}
	if x1 == x2 {
		count++
		res[x1][y1] = count
	}
	return res
}
