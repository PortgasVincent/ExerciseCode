package main

func main() {

}

func fourSum(nums []int, target int) [][]int {
	QuickSort(nums)

	var res = [][]int{}
	for k := 0; k < len(nums); k++ {
		newTarget := target - nums[k]
		for i := k + 1; i < len(nums)-2; i++ {
			left, right := i+1, len(nums)-1

			for left < right {
				sum := nums[i] + nums[left] + nums[right]
				if sum == newTarget {
					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
					for left+1 < right && nums[left+1] == nums[left] {
						left++
					}
					for right-1 > left && nums[right-1] == nums[right] {
						right--
					}
					left++
					right--
				} else if sum > newTarget {
					right--
				} else if sum < newTarget {
					left++
				}
			}

			for i+1 < len(nums) && nums[i+1] == nums[i] {
				i++
			}
		}
		for k+1 < len(nums) && nums[k+1] == nums[k] {
			k++
		}
	}
	return res
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
	return
}

func quickSort(nums []int, left, right int) {
	temp := nums[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
		for i <= p && nums[i] <= temp {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
	}
	nums[p] = temp

	if p-left > 1 {
		quickSort(nums, left, p-1)
	}
	if right-p > 1 {
		quickSort(nums, p+1, right)
	}
	return
}
