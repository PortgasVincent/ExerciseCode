package main

func main() {

}

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	return search(nums, 0, l-1, target)
}

func search(nums []int, left, right, target int) int {
	if left == right {
		if nums[left] == target {
			return left
		}
		if nums[left] < target {
			return left + 1
		}
		if nums[left] > target {
			return left
		}
	}
	if right-left == 1 {
		if nums[left] >= target {
			return left
		}
		if nums[right] >= target {
			return right
		}
		return right + 1
	}

	k := (left + right) / 2
	if nums[k] == target || (nums[k] > target && nums[k-1] < target) {
		return k
	}
	if nums[k] > target {
		return search(nums, left, k-1, target)
	}
	return search(nums, k+1, right, target)
}
