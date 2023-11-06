package main

import (
	"fmt"
)

func main() {
	nums := [][]byte{{'0', '1'}, {'0', '1'}}
	fmt.Println(maximalRectangle(nums))
}

func maximalRectangle(matrix [][]byte) int {
	lx := len(matrix)
	ly := len(matrix[0])
	left := [][]int{}
	res := [][]int{}
	for i := 0; i < lx; i++ {
		left = append(left, []int{})
		res = append(res, []int{})
		for j := 0; j < ly; j++ {
			left[i] = append(left[i], 0)
			res[i] = append(res[i], 0)
		}
	}
	for i := 0; i < lx; i++ {
		if matrix[i][0] == '0' {
			left[i][0] = 0
			continue
		}
		left[i][0] = 1
	}
	for i := 0; i < lx; i++ {
		for j := 1; j < ly; j++ {
			if matrix[i][j] == '0' {
				left[i][j] = 0
				continue
			}
			left[i][j] = left[i][j-1] + 1
		}
	}

	for j := 0; j < ly; j++ {
		stack := []int{}
		subLeft := make([]int, lx)
		subRight := make([]int, lx)
		for i := 0; i < lx; i++ {
			if len(stack) == 0 {
				stack = []int{i}
				subLeft[i] = -1
				continue
			}
			k := len(stack) - 1
			for ; k >= 0; k-- {
				if left[stack[k]][j] >= left[i][j] {
					subRight[stack[k]] = i - 1
				} else {
					break
				}
			}
			if k < 0 {
				stack = []int{i}
				subLeft[i] = -1
				continue
			}
			subLeft[i] = stack[k]
			stack = stack[:k+1]
			stack = append(stack, i)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			subRight[stack[i]] = lx - 1
		}
		for i := 0; i < lx; i++ {
			res[i][j] = (subRight[i] - subLeft[i]) * left[i][j]
		}
	}
	var max int = -1
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if res[i][j] > max {
				max = res[i][j]
			}
		}
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	lx := len(matrix)
	ly := len(matrix[0])
	left, res := make([][]int, lx), make([][]int, lx)
	for i := 0; i < lx; i++ {
		left[i] = make([]int, ly)
		res[i] = make([]int, ly)
	}
	for i := 0; i < lx; i++ {
		if matrix[i][0] == '1' {
			left[i][0] = 1
		}
	}
	for i := 0; i < lx; i++ {
		for j := 1; j < ly; j++ {
			if matrix[i][j] == '0' {
				left[i][j] = 0
				continue
			}
			left[i][j] = left[i][j-1] + 1
		}
	}
	for j := 0; j < ly; j++ {
		subLeft := make([]int, lx)
		subRight := make([]int, lx)
		stack := []int{}
		for i := 0; i < lx; i++ {
			if len(stack) == 0 {
				subLeft[i] = -1
				stack = append(stack, i)
				continue
			}
			k := len(stack) - 1
			for ; k >= 0; k-- {
				if left[i][j] >= left[stack[k]][j] {
					break
				}
				subRight[k] = i
			}
			if k < 0 {
				subLeft[i] = -1
				stack = []int{i}
				continue
			}
			subLeft[i] = k
			stack = stack[:k+1]
			stack = append(stack, i)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			subRight[stack[i]] = lx - 1
		}

		for i := 0; i < lx; i++ {
			res[i][j] = (subRight[i] - subLeft[i]) * left[i][j]
		}
	}

	max := -1
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if res[i][j] > max {
				max = res[i][j]
			}
		}
	}
	return max
}
