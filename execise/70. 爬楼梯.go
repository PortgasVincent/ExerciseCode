package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(3))
}

var record = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if x, ok := record[n]; ok {
		return x
	}
	sum := climbStairs(n-1) + climbStairs(n-2)
	record[n] = sum
	return sum
}
