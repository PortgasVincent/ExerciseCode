package main

import (
	"fmt"
)

func main() {

	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func subsets(nums []int) [][]int {
	l := len(nums)
	res := [][]int{}
	if l == 1 {
		res = append(res, []int{})
		res = append(res, []int{nums[0]})
		return res
	}

	sub := subsets(nums[:l-1])
	res = append(res, sub...)
	for _, v := range sub {
		new := []int{nums[l-1]}
		new = append(new, v...)
		res = append(res, new)
	}
	return res
}
