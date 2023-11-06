package main

import "fmt"

func main() {
	nums := []int{1, 4, 6, 6, 3, 2, 4}
	QuickSort(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		var temp int = 101
		var index int = -1
		for j := i + 1; j < l; j++ {
			if nums[j] > nums[i] && nums[j] < temp {
				temp = nums[j]
				index = j
			}
		}
		if index != -1 {
			nums[index] = nums[i]
			nums[i] = temp
			QuickSort(nums[i+1:])
			return
		}
	}

	QuickSort(nums)
	return
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
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
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
