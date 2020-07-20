package sort

import (
	"fmt"
)

func main() {
	arr := []int{4, 6, 3, 4, 9, 6}
	DichotomySort(arr)
	fmt.Println(arr)
}

func DichotomySort(arr []int){
	if len(arr) <= 1{
		return
	}
	for i := 0; i < len(arr); i++{
		left, right := 0, i
		temp := arr[i]
		var mid int
		for left <= right{
			mid = (left + right)/2
			if arr[mid] > temp{
				right = mid - 1
				continue
			}
			left = mid + 1
		}
		for j := i; j > mid; j--{
			arr[j] = arr[j-1]
			// has run arr[mid+1] = arr[mid]
		}

		// arr[mid+1]和arr[mid]都等于arr[mid] 只需根据情况将temp放到mid或者mid+1这两个位置即可
		if arr[mid] < temp{
			arr[mid+1] = temp
		} else if arr[mid] > temp {
			arr[mid] = temp
		}
		fmt.Printf("arr:%v\n", arr)
	}
}