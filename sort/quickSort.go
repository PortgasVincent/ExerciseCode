package sort

import (
	"fmt"
)

// func quickSort(values []int, left, right int) {
// 	i, j := left, right
// 	p := left
// 	temp := values[p]
// 	for i <= j {
// 		for j >= p && values[j] >= temp {
// 			j--
// 		}
// 		if j >= p {
// 			values[p] = values[j]
// 			p = j
// 		}
// 		for i <= p && values[i] <= values[p] {
// 			i++
// 		}
// 		if i <= p {
// 			values[p] = values[i]
// 			p = i
// 		}
// 	}
// 	values[p] = temp
// 	if p-left > 1 {
// 		quickSort(values, left, p-1)
// 	}
// 	if right-p > 1 {
// 		quickSort(values, p+1, right)
// 	}
// }

// func QuickSort(values []int) {
// 	if len(values) <= 1 {
// 		return
// 	}
// 	quickSort(values, 0, len(values)-1)
// }
//
//func QuickSort(arr []int){
//	if len(arr) <= 1 {
//		return
//	}
//	quickSort(arr, 0, len(arr)-1)
//	return
//}
//
//func quickSort(arr []int, left, right int){
//	i, j := left, right
//	p := left
//	temp := arr[left]
//
//	for i <= j {
//		for p <= j && arr[j] >= temp{
//			j--
//		}
//		if p <= j {
//			arr[p] = arr[j]
//			p = j
//		}
//		for i <= p && arr[i] <= temp {
//			i++
//		}
//		if i <= p {
//			arr[p] = arr[i]
//			p = i
//		}
//	}
//	arr[p] = temp
//	if p - left >= 1 {
//		quickSort(arr, left, p-1)
//	}
//	if right - p >= 1{
//		quickSort(arr, p+1, right)
//	}
//}

func main() {
	arr := []int{1, 5, 8, 56, 4, 6, 3, 9, 45}
	QuickSort(arr)
	fmt.Println(arr)
}

func QuickSort(arr []int){
	if len(arr) <= 1{
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int){
	i, j := left, right
	p := left
	temp := arr[left]
	for i <= j{
		for p <= j && temp <= arr[j]{
			j--
		}
		if p <= j {
			arr[p] = arr[j]
			p = j
		}
		for p >= i && temp >= arr[i]{
			i++
		}
		if p >= i{
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp
	if p - left >= 1{
		quickSort(arr, left, p-1)
	}
	if right - p >= 1{
		quickSort(arr, p+1, right)
	}
	return
}