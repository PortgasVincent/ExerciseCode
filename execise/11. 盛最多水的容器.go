package main

func main() {

}

// // 暴力
// func maxArea(height []int) int {
// 	length := len(height)
// 	if length < 2 {
// 		return 0
// 	}
// 	var res int
// 	for i := 0; i < length; i++ {
// 		for j := i; j < length; j++ {
// 			area := (j-i) * min(height[i], height[j])
// 			if area > res {
// 				res = area
// 			}
// 		}
// 	}
// 	return res
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Recommand
func maxArea(height []int) int {
	length := len(height)

	var res int
	for i, j := 0, length-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		if area > res {
			res = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
