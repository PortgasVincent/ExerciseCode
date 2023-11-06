package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums), nums)
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	i, j, k := 0, 0, 1
	var last int = -1
	for k < l {
		if nums[k] != nums[k-1] {
			nums[i] = nums[k-1]
			i++

			j = k
			k++
			continue
		}
		if last != j {
			nums[i] = nums[j]
			i = i + 1
			last = j
		}

		k++
	}
	nums[i] = nums[l-1]
	i++
	return i
}
