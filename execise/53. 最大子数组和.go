package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 超时
// func maxSubArray(nums []int) int {
// 	var max int = math.MinInt32
// 	var end int = len(nums) - 1
// 	for i := 0; i < len(nums); i++ {
// 		sum := 0
// 		if end < i {
// 			end = len(nums) - 1
// 		}
// 		var temp int
// 		for j := i; j <= end; j++ {
// 			sum += nums[j]
// 			if sum > max {
// 				max = sum
// 				temp = j
// 			}
// 		}
// 		end = temp
// 	}
// 	return max
// }

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	var max, sum = nums[0], nums[0]
	for i := 1; i < l; i++ {
		if sum > 0 {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// 分治
type Status struct {
	L, R, Sum, SubSum int // 从左边起最大值，从右边起最大值，全部的和，子数组最大值
}

func maxSubArray(nums []int) int {
	status := GetStatus(nums, 0, len(nums)-1)
	return status.SubSum
}

func GetStatus(nums []int, left, right int) Status {
	if left == right {
		return Status{
			L:      nums[left],
			R:      nums[left],
			Sum:    nums[left],
			SubSum: nums[left],
		}
	}
	m := (left + right) / 2
	lStatus := GetStatus(nums, left, m)
	rStatus := GetStatus(nums, m+1, right)
	return merge(lStatus, rStatus)
}

func merge(lStatus, rStatus Status) Status {
	l := max(lStatus.L, lStatus.Sum+rStatus.L)
	r := max(rStatus.R, rStatus.Sum+lStatus.R)
	sum := lStatus.Sum + rStatus.Sum
	subSum := max(lStatus.R+rStatus.L, max(lStatus.SubSum, rStatus.SubSum))
	return Status{
		L:      l,
		R:      r,
		Sum:    sum,
		SubSum: subSum,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
