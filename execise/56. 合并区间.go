package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		j := len(res) - 1

		// intervals 排过序，所以start只能>= res[j][0]
		if start > res[j][1] {
			res = append(res, intervals[i])
			continue
		}
		if end > res[j][1] {
			res[j][1] = end
		}

	}
	return res
}
