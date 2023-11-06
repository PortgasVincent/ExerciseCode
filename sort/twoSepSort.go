package main

import (
	"fmt"
)

// func DichotomySort(arr []int) {
// 	if len(arr) <= 1 {
// 		return
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		left, right := 0, i
// 		temp := arr[i]
// 		var mid int
// 		for left <= right {
// 			mid = (left + right) / 2
// 			if arr[mid] > temp {
// 				right = mid - 1
// 				continue
// 			}
// 			left = mid + 1
// 		}
// 		for j := i; j > mid; j-- {
// 			arr[j] = arr[j-1]
// 			// has run arr[mid+1] = arr[mid]
// 		}

// 		// arr[mid+1]和arr[mid]都等于arr[mid] 只需根据情况将temp放到mid或者mid+1这两个位置即可
// 		if arr[mid] < temp {
// 			arr[mid+1] = temp
// 		} else if arr[mid] > temp {
// 			arr[mid] = temp
// 		}
// 		fmt.Printf("arr:%v\n", arr)
// 	}
// }

func main() {
	arr := []int{3, 2, 23, 45, 4, 6, 3, 4, 9, 5, 3, 2, 1, 1, 6, 6}
	DichotomySort(arr)
	fmt.Println(arr)
}

func DichotomySort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		left, right := 0, i
		temp := arr[i]
		var mid int

		for left <= right {
			mid = (left + right) / 2
			if arr[mid] > temp {
				right = mid - 1
				continue
			}
			left = mid + 1
		}

		for j := i; j > mid; j-- {
			arr[j] = arr[j-1]
		}

		if arr[mid] > temp {
			arr[mid] = temp
		} else if arr[mid] < temp {
			arr[mid+1] = temp
		}
	}
	return
}
