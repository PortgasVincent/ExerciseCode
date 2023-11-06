package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// // 递归
// func trap(height []int) int {
// 	if len(height) <= 2 {
// 		return 0
// 	}
// 	if height[0] <= height[1] {
// 		return trap(height[1:])
// 	}
// 	var sum, h1, h2, temp, i int
// 	tempSum := []int{0, height[0] - height[1]}
// 	h1 = height[0]
// 	h2 = height[1]
// 	temp = 1
// 	for i = 2; i < len(height); i++ {
// 		if height[i] > h2 {
// 			h2 = height[i]
// 			temp = i
// 		}
// 		if height[i] >= h1 {
// 			for _, v := range tempSum {
// 				sum += v
// 			}
// 			break
// 		}
// 		tempSum = append(tempSum, h1-height[i])
// 	}
// 	if i < len(height) {
// 		return sum + trap(height[i:])
// 	}
// 	for i := 1; i < temp; i++ {
// 		sum += tempSum[i] - (h1 - h2)
// 	}
// 	return sum + trap(height[temp:])
// }

// // 未优化动态规划   时间O(n) 空间O(n)
// func trap(height []int) int {
// 	l := len(height)
// 	leftHeight := make([]int, l)
// 	rightHeight := make([]int, l)
// 	for i := 1; i < l; i++ {
// 		leftHeight[i] = max(leftHeight[i-1], height[i-1])
// 	}
// 	for i := l - 2; i >= 0; i-- {
// 		rightHeight[i] = max(rightHeight[i+1], height[i+1])
// 	}
// 	sum := 0
// 	for i := 0; i < l; i++ {
// 		sum += max(0, min(leftHeight[i], rightHeight[i])-height[i])
// 	}
// 	return sum
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 优化后的动态规划   时间O(n) 空间O(1)
func trap(height []int) int {
	l := len(height)

	sum := 0
	i, j, leftHeight, rightHeight := 1, l-2, height[0], height[l-1]
	for i <= j {
		if leftHeight <= rightHeight {
			sum += max(0, leftHeight-height[i])
			leftHeight = max(leftHeight, height[i])
			i++
		} else {
			sum += max(0, rightHeight-height[j])
			rightHeight = max(rightHeight, height[j])
			j--
		}
	}
	return sum
}
