package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	arr = Push[int](4, arr)
	fmt.Println(arr)
}

func Push[T any](val T, arr []T) []T {
	arr = append(arr, val)
	return arr
}
