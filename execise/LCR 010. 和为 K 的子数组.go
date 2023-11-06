package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	l := len(nums)
	sumMap := map[int]int{0: 1}
	sum := 0
	count := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		count += sumMap[sum-k]
		sumMap[sum] += 1
	}
	fmt.Println(sumMap)
	return count
}
