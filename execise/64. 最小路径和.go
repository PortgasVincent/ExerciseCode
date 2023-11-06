package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func minPathSum(grid [][]int) int {
	minMap := map[string]int{}
	m := len(grid)
	n := len(grid[0])
	return getPath(m, n, minMap, grid)
}

func getPath(m, n int, minMap map[string]int, grid [][]int) int {
	if min, ok := minMap[getKey(m, n)]; ok {
		return min
	}
	if m == 1 && n == 1 {
		return grid[0][0]
	}
	if m < 1 || n < 1 {
		return math.MaxInt32
	}
	min := grid[m-1][n-1] + min(getPath(m-1, n, minMap, grid), getPath(m, n-1, minMap, grid))
	minMap[getKey(m, n)] = min
	return min
}

func getKey(m, n int) string {
	return strconv.Itoa(m) + "," + strconv.Itoa(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
