package leetcode

import "fmt"

func main(){
	fmt.Println(numIdenticalPairs([]int{1,2,3,1,1,3}))
}

func numIdenticalPairs(nums []int) int {
	counts := 0
	sorts := make([]int, len(nums))
	for i := len(nums)-1; i >= 0; i-- {
		counts += find1(sorts, nums[i], len(nums) - i - 1)
	}
	//fmt.Printf("sorts:%v\n", sorts)
	return counts
}

func find1(sorts []int, val, l int) int {
	if l == 0{
		sorts[0] = val
		return 0
	}
	index := check1(sorts, val, 0, l-1)
	sum := 0
	for i := l; i > index; i--{
		sorts[i] = sorts[i-1]
		if sorts[i-1] == val{
			sum += 1
			fmt.Printf("val: %v, other: %v\n", val, sorts[i-1])
		}
	}
	sorts[index] = val
	fmt.Printf("sorts:%v\n", sorts)
	return sum
}

func check1(sorts []int, val int, left, right int) int {
	if left ==  right {
		if sorts[left] < val{
			return left + 1
		}
		return left
	}
	mid := (left + right)/2
	if sorts[mid] >= val {
		if mid > left {
			return check1(sorts, val, left, mid-1)
		}
		return check1(sorts, val, left, mid)
	}
	if mid < right {
		return check1(sorts, val, mid+1, right)
	}
	return check1(sorts, val, mid, right)
}
