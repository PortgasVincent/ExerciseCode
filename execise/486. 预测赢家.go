package main

func main() {

}

func predictTheWinner(nums []int) bool {
	l := len(nums)
	return pickBig(nums, 0, 0, 0, l-1, 0)
}

func pickBig(nums []int, scoreA, scoreB, left, right, count int) bool {
	if left == right {
		if count == 0 {
			if scoreA+nums[left] >= scoreB {
				return true
			}
			return false
		}
		//count == 1
		if scoreB+nums[left] > scoreA {
			return false
		}
		return true
	}
	if count == 0 {
		return pickBig(nums, scoreA+nums[left], scoreB, left+1, right, 1) || pickBig(nums, scoreA+nums[right], scoreB, left, right-1, 1)
	}
	//count == 1
	return pickBig(nums, scoreA, scoreB+nums[left], left+1, right, 0) && pickBig(nums, scoreA, scoreB+nums[right], left, right-1, 0)
}
