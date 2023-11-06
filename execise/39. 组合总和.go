package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}

func combinationSum(candidates []int, target int) [][]int {
	QuickSort(candidates)

	return comnbina(candidates, target, []int{})
}

func comnbina(candidates []int, target int, prefix []int) [][]int {
	if target == 0 {
		return [][]int{prefix}
	}
	if target < candidates[0] {
		return nil
	}
	var res = [][]int{}
	for i := 0; i < len(candidates); i++ {
		v := candidates[i]
		if v > target {
			break
		}
		if v == target {
			res = append(res, append(prefix, v))
			break
		}
		newPrefix := append(clone(prefix), v)
		res = append(res, comnbina(candidates[i:], target-v, newPrefix)...)
	}
	return res
}

func clone(old []int) []int {
	var new = make([]int, len(old))
	for i, v := range old {
		new[i] = v
	}
	return new
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
