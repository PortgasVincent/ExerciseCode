package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

// // 超时
// func threeSum(nums []int) [][]int {
// 	var reverseArr = map[int][]int{}
// 	for i, v := range nums {
// 		_, ok := reverseArr[v]
// 		if ok {
// 			reverseArr[v] = append(reverseArr[v], i)
// 			continue
// 		}
// 		reverseArr[v] = []int{i}
// 	}

// 	var res = [][]int{}
// 	var repeated = map[string]bool{}
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			index := -(nums[i] + nums[j])
// 			indexArr, ok := reverseArr[index]
// 			if ok {
// 				for _, v := range indexArr {
// 					if v <= j {
// 						continue
// 					}
// 					if repeated[order(nums[i], nums[j], nums[v])] {
// 						continue
// 					}
// 					res = append(res, []int{nums[i], nums[j], nums[v]})
// 					repeated[order(nums[i], nums[j], nums[v])] = true
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func order(a, b, c int) string {
// 	if a >= b {
// 		if b >= c {
// 			return strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(c), 10)
// 		}
// 		if a >= c {
// 			return strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(c), 10) + strconv.FormatInt(int64(b), 10)
// 		}
// 		return strconv.FormatInt(int64(c), 10) + strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(b), 10)
// 	}
// 	if a >= c {
// 		return strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(c), 10)
// 	}
// 	if c >= b {
// 		return strconv.FormatInt(int64(c), 10) + strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(a), 10)
// 	}
// 	return strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(c), 10) + strconv.FormatInt(int64(a), 10)
// }

func threeSum(arr []int) [][]int {
	QuickSort(arr)

	var res [][]int
	for i := 0; i < len(arr)-2; i++ {
		if arr[i] > 0 {
			break
		}
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left := i + 1
		right := len(arr) - 1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				res = append(res, []int{arr[i], arr[left], arr[right]})
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			}
			if sum > 0 {
				right--
			}
			if sum < 0 {
				left++
			}
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
