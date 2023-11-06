package main

import (
	"fmt"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}

func permuteUnique(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	if l == 1 {
		return [][]int{{nums[0]}}
	}

	var res = [][]int{}
	var m = map[int]bool{}
	for i := 0; i < l; i++ {
		if m[nums[i]] {
			continue
		}
		arr := append([]int{}, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		subRes := permuteUnique(arr)
		for _, v := range subRes {
			res = append(res, append([]int{nums[i]}, v...))
		}

		m[nums[i]] = true
	}
	return res
}
