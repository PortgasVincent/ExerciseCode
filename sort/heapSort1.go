package main

import (
	"fmt"
)

func main(){
	a := []int{ 3,2,45,1,5, }
	HeapSort(a)
	fmt.Println(a)   
}

func HeapSort(arr []int){
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := l/2-1;i>=0;i--{
		heapSort(arr, i, l)
	}
	l--
	for l>0 {
		arr[0], arr[l] = exchange(arr[0], arr[l])
		heapSort(arr, 0, l)
		l--
	}
	
}

func heapSort(arr []int, i, l int){
	temp := arr[i]
	for k := 2*i+1; k < l; k = 2*k+1 {
		if k+1 < l && arr[k+1] > arr[k] {
			k = k+1
		}
		if arr[k] > arr[i] {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}

func exchange(a, b int) (int, int){
	return b, a
}