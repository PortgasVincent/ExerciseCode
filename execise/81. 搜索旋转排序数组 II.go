package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums), nums)
}

func search(nums []int, target int) bool {
	l := len(nums)
	return find(nums, 0, l-1, target)
}

func find(nums []int, left, right, target int) bool {
	if right <= left {
		return (left < len(nums) && nums[left] == target) || (right >= 0 && nums[right] == target)
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return true
	}
	if nums[mid] > target {
		if nums[mid] >= nums[left] {
			return find(nums, mid+1, right, target) || find(nums, left, mid-1, target)
		}
		return find(nums, left, mid-1, target)
	}
	if nums[mid] > nums[left] {
		return find(nums, mid+1, right, target)
	}
	return find(nums, mid+1, right, target) || find(nums, left, mid-1, target)
}
