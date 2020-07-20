package main

import "fmt"

func main(){
	// fmt.Println(findNumberIn2DArray([][]int{
	// 	{1,4,7,11,15},
	// 	{2,5,8,12,19},
	// 	{3,6,9,16,22},
	// 	{10,13,14,17,24},
	// 	{18,21,23,26,30}}, 5))
	fmt.Println(
		findNumberIn2DArray(
			[][]int{
				{1,4},
				{2,5}},2))
}


// func findNumberIn2DArray(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}
// 	x := len(matrix[0]) - 1
// 	y := len(matrix) - 1
// 	originX, originY := x, y
// 	i, j := 0, 0
// 	for ; i <= x || j <= y;{
// 		fmt.Println(i, j, x, y)
// 		if matrix[j][i] > target || matrix[y][x] < target {
// 			return false
// 		}
// 		if matrix[j][i] == target || matrix[y][x] == target || matrix[y][i] == target || matrix[j][x] == target {
// 			return true
// 		}
// 		if y - j <= 1 && x - i <= 1 {
// 			break
// 		}
// 		if matrix[j+(y-j)/2][i+(x-i)/2] == target {
// 			return true
// 		}
// 		if matrix[j+(y-j)/2][i+(x-i)/2] > target {
// 			return findNumberIn2DArray(matrix[j:j+(y-j)/2][i:i+(x-i)/2], target) ||
// 			 findNumberIn2DArray(matrix[j:j+(y-j)/2][i+(x-i)/2:x], target) ||
// 			 findNumberIn2DArray(matrix[j+(y-j)/2:y][i:i+(x-i)/2], target)
// 		} else {
// 			return findNumberIn2DArray(matrix[j+(y-j)/2:y][i+(x-i)/2:x], target) ||
// 			 findNumberIn2DArray(matrix[j:j+(y-j)/2][i+(x-i)/2:x], target) ||
// 			 findNumberIn2DArray(matrix[j+(y-j)/2:y][i:i+(x-i)/2], target)
// 		}
// 	}
// 	var temp = i
// 	for ;i<=originX;i++{
// 		if matrix[j][i] == target {
// 			return true
// 		}
// 	}
// 	for ;j<=originY;j++{
// 		if matrix[j][temp] == target {
// 			return true
// 		}
// 	}
// 	return false
// }

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix[0]);i++{
		for j := 0; j<len(matrix);j++{
			if matrix[j][i] == target {
				return true
			}
		}
	}
	return false
}