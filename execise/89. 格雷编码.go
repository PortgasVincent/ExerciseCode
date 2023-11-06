package main

import (
	"fmt"
)

func main() {
	nums := [][]byte{{'0', '1'}, {'0', '1'}}
	fmt.Println(maximalRectangle(nums))
}

func grayCode(n int) []int {
	res := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head <<= 1
	}
	return res
}
