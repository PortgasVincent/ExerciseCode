package main

import "fmt"

func main() {
	fmt.Println(combinationSum2([]int{7, 3, 2}, 18))
}

func combinationSum2(candidates []int, target int) [][]int {
	QuickSort(candidates)
	return combine(candidates, target, []int{})
}

func combine(candidates []int, target int, prefix []int) [][]int {
	var res = [][]int{}
	var last int = 51
	for i, v := range candidates {
		if last == v {
			continue
		}
		last = v
		if v > target {
			continue
		}
		if v == target {
			res = append(res, append(copy(prefix), v))
		}
		if i == len(candidates)-1 {
			break
		}
		res = append(res, combine(candidates[i+1:], target-v, append(copy(prefix), v))...)
	}
	return res
}

func copy(old []int) []int {
	new := make([]int, len(old))
	for i, v := range old {
		new[i] = v
	}
	return new
}

func newArr(arr []int, x int) []int {
	res := []int{}
	for i, v := range arr {
		if i != x {
			res = append(res, v)
		}
	}
	return res
}

func QuickSort(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	quickSort(nums, 0, l-1)
	return
}
func quickSort(nums []int, left, right int) {
	i, j := left, right
	p := left
	temp := nums[p]

	for i <= j {
		for p <= j && nums[j] >= temp {
			j--
		}
		if p <= j {
			nums[p] = nums[j]
			p = j
		}
		for i <= p && nums[i] <= temp {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = temp

	if p-left > 1 {
		quickSort(nums, left, p-1)
	}
	if right-p > 1 {
		quickSort(nums, p+1, right)
	}
	return
}
