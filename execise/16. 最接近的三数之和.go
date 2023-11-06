package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}

func threeSumClosest(nums []int, target int) int {
	QuickSort(nums)
	var diff int = math.MaxInt
	var res int
	for i := 0; i < len(nums)-1; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if int(math.Abs(float64(sum-target))) < diff {
				diff = int(math.Abs(float64(sum - target)))
				res = sum
			}
			if sum-target > 0 {
				right--
				continue
			}
			left++
		}
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
	temp := nums[left]
	p := left
	i, j := left, right
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
