package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("asdfgfdsa"))
}

// func longestPalindrome(s string) string {
// 	a := []rune(s)
// 	l := len(a)
// 	var temp string
// 	for i := 0; i < l; i++{
// 		res, length := One(a, i)
// 		if length > len(temp) {
// 			temp = res
// 		}
// 		res, length = Two(a, i)
// 		if length > len(temp) {
// 			temp = res
// 		}
// 	}
// 	return temp
// }

// func One(a []rune, i int)(res string, l int){
// 	res = string(a[i])
// 	l = 1
// 	for j := 0; i - j >= 0 && i + j < len(a); j++{
// 		if a[i-j] == a[i+j]{
// 			res = string(a[i-j:i+j+1])
// 			l = len(res)
// 		} else {
// 			return
// 		}
// 	}
// 	return res, l
// }

// func Two(a []rune, i int)(res string, l int){
// 	for j := 0; i - j >= 0 && i + j + 1 < len(a); j++{
// 		if a[i-j] == a[i+j+1]{
// 			res = string(a[i-j:i+j+1+1])
// 			l = len(res)
// 		} else {
// 			return
// 		}
// 	}
// 	return
// }

// 暴力破解，O(n^3)
// func longestPalindrome(s string) string {
// 	arr := []rune(s)
// 	var res []rune
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j <= len(arr); j++ {
// 			if isPalindrome(string(arr[i:j])) && len(arr[i:j]) > len(res) {
// 				res = arr[i:j]
// 			}
// 		}
// 	}
// 	return string(res)
// }

// func isPalindrome(s string) bool {
// 	if s == reverse(s) {
// 		return true
// 	}
// 	return false
// }

// func reverse(s string) string {
// 	arr := []rune(s)
// 	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}
// 	return string(arr)
// }

// // 动态规划
// func longestPalindrome(s string) string {
// 	arr := []rune(s)
// 	length := len(arr)
// 	if length < 2 {
// 		return s
// 	}

// 	var maxLen, begin int
// 	maxLen = 1
// 	m := [][]bool{}
// 	for i := 0; i < length; i++ {
// 		m = append(m, []bool{})
// 		for j := 0; j < length; j++ {
// 			m[i] = append(m[i], false)
// 		}
// 	}

// 	for i := 0; i < length; i++ {
// 		m[i][i] = true
// 	}

// 	// 按l便利，现填充对角线附近的值，m[i+1][j-1]会是对角线上，也会有值
// 	for l := 1; l < length; l++ {
// 		for i := 0; i < length; i++ {
// 			j := i + l

// 			if j >= length {
// 				break
// 			}

// 			if arr[i] != arr[j] {
// 				m[i][j] = false
// 			} else {
// 				if l < 3 {
// 					m[i][j] = true
// 				} else {
// 					m[i][j] = m[i+1][j-1]
// 				}
// 			}

// 			if m[i][j] && j-i+1 > maxLen {
// 				maxLen = j - i + 1
// 				begin = i
// 			}
// 		}
// 	}

// 	return string(arr[begin : begin+maxLen])
// }

// 中心扩散法  推荐！！！
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	arr := []rune(s)
	var start, end int
	for i := 0; i < len(arr); i++ {
		left1, right1 := expandAround(arr, i, i)
		left2, right2 := expandAround(arr, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return string(arr[start : end+1])
}

func expandAround(arr []rune, left, right int) (int, int) {
	for ; left >= 0 && right < len(arr) && arr[left] == arr[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
