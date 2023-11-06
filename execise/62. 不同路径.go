package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	exists := map[string]int{}
	return getPath(m, n, exists)
}

func getPath(m, n int, record map[string]int) int {
	if count, ok := record[parseKey(m, n)]; ok {
		return count
	}

	if m == 1 || n == 1 {
		return 1
	}
	sum := getPath(m-1, n, record) + getPath(m, n-1, record)
	record[parseKey(m, n)] = sum
	return sum
}

func parseKey(m, n int) string {
	return strconv.Itoa(m) + "," + strconv.Itoa(n)
}
