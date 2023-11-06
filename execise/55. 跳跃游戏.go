package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func canJump(nums []int) bool {
	l := len(nums) - 1
	start := 0
	end := 0
	for end < l {
		var max = 0
		for i := start; i <= end; i++ {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
		if max >= l {
			return true
		}
		if max == end {
			return false
		}
		start = end + 1
		end = max
	}

	return true
}
