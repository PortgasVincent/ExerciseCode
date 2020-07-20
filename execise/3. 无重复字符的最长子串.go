package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	a := []rune(s)
	m := map[rune]int{}
	temp := 0
	maxL := 0
	startIndex := 0
	for i := 0; i < len(a); i++{
		index, ok := m[a[i]]
		if ok && index >= startIndex {
			repeatLen := index + 1 - startIndex
			temp = temp - repeatLen
			startIndex = index + 1
		}
		temp++

		if temp > maxL {
			maxL = temp
		}
		m[a[i]] = i
	}
	return maxL
}

func main(){
	fmt.Printf("max length: %v", lengthOfLongestSubstring("pwwkew"))
}