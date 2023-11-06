package main

import (
	"fmt"
)

func swap(a int, b int) (int, int) {
	return b, a
}

// func HeapSort(arr []int){
// 	length := len(arr)
// 	if length <= 1{
// 		return
// 	}
// 	for i := length/2-1;i >= 0;i--{
// 		adjust(arr, i, length)
// 	}
// 	for  i:= length -1; i >= 0; i--{
// 		arr[0], arr[i] = swap(arr[0], arr[i])
// 		adjust(arr, 0, i)
// 	}
// }

// func adjust(arr []int, i, l int){
// 	temp := arr[i]
// 	for k := 2*i+1; k < l; k = 2*k+1{
// 		if k+1 < l && arr[k+1]>arr[k]{
// 			k += 1
// 		}
// 		if arr[k]>temp{
// 			arr[i] = arr[k]
// 			i = k
// 		}

// 	}
// 	arr[i] = temp
// }

func main() {
	arr := []int{1, 5, 5, 6, 8, 56, 4, 6, 3, 9, 45}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}

	for i := l/2 - 1; i >= 0; i-- {
		adjust(arr, i, l-1)
	}
	for i := l - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust(arr, 0, i-1)
	}
	return
}

func adjust(arr []int, left, right int) {
	temp := arr[left]
	for k := 2*left + 1; k <= right; k = 2*k + 1 {
		if k+1 <= right && arr[k+1] > arr[k] {
			k++
		}
		if arr[k] > temp {
			arr[left] = arr[k]
			left = k
		}
	}
	arr[left] = temp
	return
}
