package main

import "fmt"

func main() {
	arr := []int{3, 2, 23, 45, 4, 6, 3, 4, 9, 5, 3, 2, 1, 1, 6, 6}
	BubbleSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j+1 < length-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return
}
