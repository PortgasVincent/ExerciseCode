package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	m := map[string]bool{}
	dupRes := set(nums)
	for i := 0; i < len(dupRes); i++ {
		if !m[getKey(dupRes[i])] {
			m[getKey(dupRes[i])] = true
			res = append(res, dupRes[i])
		}
	}
	return res
}

func set(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}

	var res = [][]int{}
	sub := set(nums[1:])
	for i := 0; i < len(sub); i++ {
		res = append(res, sub[i])
		new := []int{nums[0]}
		res = append(res, append(new, sub[i]...))
	}
	return res
}

func getKey(nums []int) string {
	if len(nums) == 0 {
		return "len0"
	}
	QuickSort(nums)
	var res string = strconv.Itoa(nums[0])
	for i := 1; i < len(nums); i++ {
		res += "," + strconv.Itoa(nums[i])
	}
	return res
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
	return
}

func quickSort(nums []int, left, right int) {
	i, j := left, right
	p := left
	temp := nums[left]
	for i <= j {
		for p <= j && nums[j] >= temp {
			j--
		}
		if p <= j {
			nums[p] = nums[j]
			p = j
		}
		for p >= i && nums[i] <= temp {
			i++
		}
		if p >= i {
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
