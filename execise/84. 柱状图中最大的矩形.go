package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(nums))
}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	var stack = []int{}
	left, right := make([]int, l), make([]int, l)

	// 找左边界值
	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			left[i] = -1
			continue
		}
		j := len(stack) - 1
		for ; j >= 0 && heights[stack[j]] >= heights[i]; j-- {
		}
		if j < 0 {
			stack = []int{i}
			left[i] = -1
			continue
		}
		left[i] = stack[j]
		stack = stack[0 : j+1]
		stack = append(stack, i)
	}

	//找右边界值
	stack = []int{}
	for i := l - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = append(stack, i)
			right[i] = l - 1
			continue
		}
		j := len(stack) - 1
		for ; j >= 0 && heights[stack[j]] >= heights[i]; j-- {
		}
		if j < 0 {
			stack = []int{i}
			right[i] = l - 1
			continue
		}
		right[i] = stack[j] - 1
		stack = stack[0 : j+1]
		stack = append(stack, i)
	}

	var max int = -1
	for i := 0; i < l; i++ {
		res := (right[i] - left[i]) * heights[i]
		if res > max {
			max = res
		}
	}
	return max
}
