package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []int{1, 2, 3}
	arr = Push[int](4, arr)
	fmt.Println(arr)

	arr1 := []int{0, 1, 2, 3}
	fmt.Println(arr1[4:])
	fmt.Println(strings.Contains("", ""))

	fmt.Println(byte(0), byte(0+'0'), '1', '1'-'0')

	m := map[string]int{"aaa": 0}
	_, ok := m["aaa"]
	fmt.Println(ok)
}

func Push[T any](val T, arr []T) []T {
	arr = append(arr, val)
	return arr
}
