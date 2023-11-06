package main

import (
	"fmt"
)

func main() {

	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func combine(n int, k int) [][]int {
	var res = [][]int{}
	if k == 1 {
		for i := 1; i <= n; i++ {
			res = append(res, []int{i})
		}
		return res
	}
	if k == n {
		temp := []int{}
		for i := 1; i <= n; i++ {
			temp = append(temp, i)
		}
		res = append(res, temp)
		return res
	}
	res = append(res, combine(n-1, k)...)
	for _, v := range combine(n-1, k-1) {
		res = append(res, append(v, n))
	}

	return res
}
