package leetcode

import "fmt"

func main(){
	fmt.Printf("%v", countSmaller([]int{5,2,6,1}))
}

func countSmaller(nums []int) []int {
	counts := make([]int, len(nums))
	sorts := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i--{
		counts[i] = find(sorts, nums[i], len(nums)-i-1)
	}
	return counts
}

func find(sorts []int, val ,l int) int {
	if l == 0 {
		sorts[0] = val
		return 0
	}
	res := divide(sorts, val, 0, l-1)
	for i := l; i > res; i--{
		sorts[i] = sorts[i-1]
	}
	sorts[res] = val
	return res
}

func divide(sorts []int, val int, left, right int)int{
	if left == right {
		if sorts[left] >= val{
			return left
		}
		return left + 1
	}
	mid := (left+right)/2
	if sorts[mid] >= val{
		if mid == left {
			return left
		}
		return divide(sorts, val, left, mid-1)
	}
	return divide(sorts, val, mid+1, right)
}

