package main

import (
	"fmt"
)

// func QuickSort(arr []int){
// 	if len(arr) <= 1{
// 		return
// 	}
// 	quickSort(arr, 0, len(arr)-1)
// }

// func quickSort(arr []int, left, right int){
// 	i, j := left, right
// 	p := left
// 	temp := arr[left]
//  // 必须用等号，否则数组有相同的值会死循环
// 	for i <= j{
// 		for p <= j && temp <= arr[j]{
// 			j--
// 		}
// 		if p <= j {
// 			arr[p] = arr[j]
// 			p = j
// 		}
// 		for p >= i && temp >= arr[i]{
// 			i++
// 		}
// 		if p >= i{
// 			arr[p] = arr[i]
// 			p = i
// 		}
// 	}
// 	arr[p] = temp
// 	if p - left >= 1{
// 		quickSort(arr, left, p-1)
// 	}
// 	if right - p >= 1{
// 		quickSort(arr, p+1, right)
// 	}
// 	return
// }

func main() {
	arr := []int{1, 5, 5, 6, 8, 56, 4, 6, 3, 9, 45}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	quickSort(arr, 0, l-1)
}

func quickSort(arr []int, left, right int) {
	i, j := left, right
	p := left
	temp := arr[left]

	for i <= j {
		for p <= j && arr[j] >= temp {
			j--
		}
		if p <= j {
			arr[p] = arr[j]
			p = j
		}

		for p >= i && arr[i] <= temp {
			i++
		}
		if p >= i {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp

	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
	return
}
