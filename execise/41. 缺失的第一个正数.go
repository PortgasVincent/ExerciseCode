package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}

func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] <= 0 {
			nums[i] = l + 1
		}
	}
	fmt.Println(nums)
	for i := 0; i < l; i++ {
		index := int(math.Abs(float64(nums[i])))
		fmt.Println(index)
		if index <= l && nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return l + 1
}
