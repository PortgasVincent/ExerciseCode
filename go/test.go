package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{-3, -2, -1, 0}
	fmt.Println(Find(arr))
}

func Find(arr []int) int {
	l := len(arr)
	if l == 0 {
		return 0
	}
	if l == 1 || arr[0] >= 0 {
		return arr[0]
	}
	return find(arr, 0, l-1)
}

func find(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	if right-left == 1 {
		return min(arr[left], arr[right])
	}
	mid := (left + right) / 2
	if arr[mid] >= 0 {
		return find(arr, left, mid)
	}
	return find(arr, mid, right)
}

func min(a, b int) int {
	if math.Abs(float64(a)) < math.Abs(float64(b)) {
		return a
	}
	return b
}
