package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	l := len(words)
	byteCount := len(words[0])
	length := l * byteCount
	var res []int
	for i := 0; i-1 < len(s)-length; i++ {
		fmt.Println(i, isCombine(s[i:i+length], words))
		if isCombine(s[i:i+length], words) {
			res = append(res, i)
		}
	}
	return res
}

func isCombine(s string, words []string) bool {
	byteCount := len(words[0])
	m := map[string]int{}
	for _, v := range words {
		m[v] = m[v] + 1
	}
	//i-1是因为i是每个字符串的第一个位置，应该是上个字符串的最后一个位置不超过就行了
	for i := 0; i-1 < len(s)-byteCount; i += byteCount {
		for _, v := range words {
			m[v] = m[v] + 1
		}
		exist, ok := m[s[i:i+byteCount]]
		// fmt.Println(i, i+byteCount, s[i:i+byteCount], exist, ok)
		if exist == 0 || !ok {
			return false
		}
		m[s[i:i+byteCount]] = exist - 1
	}
	return true
}
