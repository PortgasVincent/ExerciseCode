package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotate([]int{1, 2, 3}))
}

func rotate(matrix [][]int) {
	n := len(matrix)
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			for j := 0; j < n/2; j++ {
				x, y := i, j
				last := matrix[x][y]
				for k := 0; k <= 3; k++ {
					p, q := convert(x, y, n-1)
					temp := matrix[p][q]
					matrix[p][q] = last
					last = temp
					x, y = p, q
				}
			}
		}
	} else {
		for i := 0; i < n/2; i++ {
			for j := 0; j <= n/2; j++ {
				x, y := i, j
				last := matrix[x][y]
				for k := 0; k <= 3; k++ {
					p, q := convert(x, y, n-1)
					temp := matrix[p][q]
					matrix[p][q] = last
					last = temp
					x, y = p, q
				}
			}
		}
	}
	return
}

func convert(x, y, n int) (int, int) {
	return y, n - x
}
