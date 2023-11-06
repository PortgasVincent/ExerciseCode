package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func plusOne(digits []int) []int {
	l := len(digits)
	var flag bool
	for i := l - 1; i >= 0; i-- {
		temp := digits[i] + 1
		flag = temp >= 10
		if temp < 10 {
			digits[i] = temp
			return digits
		}
		digits[i] = temp % 10
	}
	if flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}
