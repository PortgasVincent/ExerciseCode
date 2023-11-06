package main

func main() {

}

func search(nums []int, target int) int {
	l := len(nums)
	return searchByScope(nums, 0, l-1, target)
}

func searchByScope(nums []int, left, right, target int) int {
	if right < left || (left == right && nums[left] != target) {
		return -1
	}
	k := (left + right) / 2
	switch {
	case target == nums[left]:
		return left
	case target == nums[k]:
		return k
	case target == nums[right]:
		return right

	}
	if nums[k] > nums[left] {
		if target > nums[k] || target < nums[left] {
			return searchByScope(nums, k+1, right, target)
		}
		return twoSepSearch(nums, left, k-1, target)
	}
	if target > nums[right] || target < nums[k] {
		return searchByScope(nums, left, k-1, target)
	}
	return twoSepSearch(nums, k+1, right, target)
}

func twoSepSearch(nums []int, left, right, target int) int {
	if right < left || (right == left && nums[left] != target) {
		return -1
	}
	k := (left + right) / 2
	if nums[k] == target {
		return k
	}
	if target > nums[k] {
		return twoSepSearch(nums, k+1, right, target)
	}
	return twoSepSearch(nums, left, k-1, target)
}
