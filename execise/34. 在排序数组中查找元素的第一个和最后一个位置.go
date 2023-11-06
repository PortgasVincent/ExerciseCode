package main

func main() {

}
func searchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	min := searchMin(nums, 0, l-1, target)
	max := searchMax(nums, 0, l-1, target)
	return []int{min, max}
}

func searchMin(nums []int, left, right, target int) int {
	if right < left || (right == left && nums[left] != target) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	if len(nums) == 2 && nums[right] == target {
		return right
	}
	k := (right + left) / 2
	if nums[k] > target || (nums[k] == target && nums[k-1] == target) {
		return searchMin(nums, left, k-1, target)
	}
	if nums[k] < target {
		return searchMin(nums, k+1, right, target)
	}
	// nums[k] == target
	return k
}

func searchMax(nums []int, left, right, target int) int {
	if right < left || (right == left && nums[left] != target) {
		return -1
	}
	if nums[right] == target {
		return right
	}
	if len(nums) == 2 && nums[left] == target {
		return left
	}
	k := (left + right) / 2
	if nums[k] > target {
		return searchMax(nums, left, k-1, target)
	}
	if nums[k] < target || (nums[k] == target && nums[k+1] == target) {
		return searchMax(nums, k+1, right, target)
	}
	// nums[k] = target
	return k
}
