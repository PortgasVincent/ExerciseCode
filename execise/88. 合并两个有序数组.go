package main

import (
	"fmt"
)

func main() {
	nums := [][]byte{{'0', '1'}, {'0', '1'}}
	fmt.Println(maximalRectangle(nums))
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m - 1; i >= 0; i-- {
		nums1[n+i] = nums1[i]
	}
	i, j, k := n, 0, 0
	for i < m+n && j < n {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	for k < m+n && i < m+n {
		nums1[k] = nums1[i]
		i++
		k++
	}
	for k < m+n && j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
	return
}
