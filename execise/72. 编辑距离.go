package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	if l1*l2 == 0 {
		return l1 + l2
	}

	dp := [][]int{}
	for i := 0; i < l1+1; i++ {
		dp = append(dp, []int{})
		for j := 0; j < l2+1; j++ {
			dp[i] = append(dp[i], math.MaxInt32)
		}
	}
	for i := 0; i < l1+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < l2+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			left := dp[i-1][j] + 1
			up := dp[i][j-1] + 1
			leftup := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftup++
			}
			dp[i][j] = min(left, min(up, leftup))
		}
	}
	return dp[l1][l2]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
