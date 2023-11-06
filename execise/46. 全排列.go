package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var res = [][]int{}
	for i := 0; i < len(nums); i++ {
		newArr := append([]int{}, nums[0:i]...)
		newArr = append(newArr, nums[i+1:]...)
		subRes := permute(newArr)
		for _, v := range subRes {
			sub := append([]int{nums[i]}, v...)
			res = append(res, sub)
		}
	}
	return res
}
