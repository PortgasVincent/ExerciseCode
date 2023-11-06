package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

// // 超时
// func jump(nums []int) int {
// 	l := len(nums)
// 	if l <= 1 {
// 		return 0
// 	}
// 	var arr = []int{}
// 	for i := 0; i < l-1; i++ {
// 		if i+nums[i] >= l-1 {
// 			arr = append(arr, i)
// 		}
// 	}
// 	var min int = math.MaxInt32
// 	for _, v := range arr {
// 		temp := jump(nums[:v+1])
// 		if temp < min {
// 			min = temp
// 		}
// 	}
// 	return min + 1
// }

func jump(nums []int) int {
	count := 0
	end := 0
	l := len(nums)
	var last int = -1
	for end < l-1 {
		temp := end
		for i := last + 1; i <= temp; i++ {
			if i+nums[i] > end {
				end = i + nums[i]
			}
		}
		fmt.Println(count, end)
		count++
		last = temp
	}
	return count
}
