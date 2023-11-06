package main

import (
	"fmt"
)

// func lengthOfLongestSubstring(s string) int {
// 	a := []rune(s)
// 	m := map[rune]int{}
// 	temp := 0
// 	maxL := 0
// 	startIndex := 0
// 	for i := 0; i < len(a); i++{
// 		index, ok := m[a[i]]
// 		if ok && index >= startIndex {
// 			repeatLen := index + 1 - startIndex
// 			temp = temp - repeatLen
// 			startIndex = index + 1
// 		}
// 		temp++

// 		if temp > maxL {
// 			maxL = temp
// 		}
// 		m[a[i]] = i
// 	}
// 	return maxL
// }

func main() {
	fmt.Printf("max length: %v", lengthOfLongestSubstring("pwwkew"))
}

// func lengthOfLongestSubstring(s string) int {
// 	var res, temp int

// 	for i := 0; i < len(s); i++ {
// 		temp = 0
// 		for j := i + 1; j <= len(s); j++ {
// 			temp++
// 			if temp > res {
// 				res = temp
// 			}
// 			if j+1 <= len(s) && strings.Contains(s[i:j], s[j:j+1]) {
// 				break
// 			}
// 		}
// 	}
// 	return res
// }

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	m := map[byte]bool{}
	i, j := 0, 0
	max := 0
	for j < l {
		if m[s[j]] {
			m[s[i]] = false
			i++
			continue
		}
		m[s[j]] = true
		j++
		if j-i > max {
			max = j - i
		}
	}
	return max
}
