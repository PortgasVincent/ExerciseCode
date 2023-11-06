package main

import "fmt"

// // return new arr
// func MergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}
// 	mid := len(arr) / 2
// 	leftSub := MergeSort(arr[:mid])
// 	rightSub := MergeSort(arr[mid:])
// 	return merge(leftSub, rightSub)
// }

// func merge(leftArr, rightArr []int) []int {
// 	lIndex, rIndex := 0, 0
// 	res := make([]int, 0)

// 	for lIndex < len(leftArr) && rIndex < len(rightArr) {
// 		if leftArr[lIndex] <= rightArr[rIndex] {
// 			res = append(res, leftArr[lIndex])
// 			lIndex++
// 		} else {
// 			res = append(res, rightArr[rIndex])
// 			rIndex++
// 		}
// 	}
// 	if lIndex < len(leftArr) {
// 		res = append(res, leftArr[lIndex:]...)
// 	}
// 	if rIndex < len(rightArr) {
// 		res = append(res, rightArr[rIndex:]...)
// 	}
// 	return res
// }

func main() {
	arr := []int{1, 5, 5, 6, 8, 56, 4, 6, 3, 9, 45}
	newArr := MergeSort(arr)
	fmt.Println(newArr)
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	lIndex, rIndex := 0, 0
	res := []int{}
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			res = append(res, left[lIndex])
			lIndex++
		} else {
			res = append(res, right[rIndex])
			rIndex++
		}
	}
	if lIndex < len(left) {
		res = append(res, left[lIndex:]...)
	}
	if rIndex < len(right) {
		res = append(res, right[rIndex:]...)
	}
	return res
}
