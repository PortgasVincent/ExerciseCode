package main

import (
	"fmt"
	"math"
)

func main() {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	fmt.Println(networkDelayTime(times, 4, 2))
}

// func networkDelayTime(times [][]int, n int, k int) int {
// 	inf := math.MaxInt / 2
// 	location := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		location[i] = make([]int, n)
// 		for index := range location[i] {
// 			location[i][index] = inf
// 		}
// 	}
// 	for _, v := range times {
// 		x := v[0] - 1
// 		y := v[1] - 1
// 		location[x][y] = v[2]
// 	}

// 	dis := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		dis[i] = inf
// 	}
// 	dis[k-1] = 0
// 	used := make([]bool, n)
// 	for i := 0; i < n; i++ {
// 		j := -1
// 		for k := 0; k < n; k++ {
// 			if !used[k] && (j == -1 || dis[k] < dis[j]) {
// 				j = k
// 			}
// 		}

// 		used[j] = true
// 		for k, v := range location[j] {
// 			dis[k] = min(dis[k], dis[j]+v)
// 			// fmt.Println(dis[k], dis[j], v)
// 			// 为避免max int做加法，应该对inf做处理。比如使其为maxInt / 2
// 		}
// 	}

// 	max := -1
// 	for i := 0; i < n; i++ {
// 		if dis[i] == inf {
// 			return -1
// 		}
// 		if dis[i] > max {
// 			max = dis[i]
// 		}
// 	}
// 	return max
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func networkDelayTime(times [][]int, n int, k int) int {
	inf := math.MaxInt / 2
	location := make([][]int, n)
	for i := 0; i < n; i++ {
		location[i] = make([]int, n)
		for j := 0; j < n; j++ {
			location[i][j] = inf
		}
	}
	for _, v := range times {
		x := v[0] - 1
		y := v[1] - 1
		location[x][y] = v[2]
	}

	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[k-1] = 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		index := -1
		for j := 0; j < n; j++ {
			if !used[j] && (index == -1 || dis[j] < dis[index]) {
				index = j
			}
		}

		used[index] = true
		for j, v := range location[index] {
			dis[j] = min(dis[j], dis[index]+v)
		}
	}
	max := -1
	for _, v := range dis {
		if v == inf {
			return -1
		}
		if v > max {
			max = v
		}
	}
	return max
}
