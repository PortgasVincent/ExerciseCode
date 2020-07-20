package main

import (
	"fmt"
)

func main(){
	a := []int{ 3,2,45,1,5, }
	QuickSort(a)
	fmt.Println(a)
}

func QuickSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	quickSort(arr, 0, l-1)
}

func quickSort(arr []int, left, right int){
	temp := arr[left]
	p := left
	i, j := left, right
	for i<j {
		if j > p && arr[j] > temp{
			j--
		}
		if j > p {
			arr[p] = arr[j]
			p = j
		}
		if i<p && arr[i]<temp {
			i++
		}
		if i<p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp
	if p - left > 1 {
		quickSort(arr, left, p-1)
	}
	if right - p > 1 {
		quickSort(arr, p+1, right)
	}
}