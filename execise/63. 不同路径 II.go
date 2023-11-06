package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	solution := map[string]int{}
	block := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				block[getKey(i+1, j+1)] = true
			}
		}
	}
	return getPath(m, n, solution, block)
}

func getPath(m, n int, solution map[string]int, block map[string]bool) int {
	if block[getKey(m, n)] {
		return 0
	}
	if count, ok := solution[getKey(m, n)]; ok {
		return count
	}
	if m < 1 || n < 1 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	sum := getPath(m-1, n, solution, block) + getPath(m, n-1, solution, block)
	solution[getKey(m, n)] = sum
	return sum
}

func getKey(m, n int) string {
	return strconv.Itoa(m) + "," + strconv.Itoa(n)
}
