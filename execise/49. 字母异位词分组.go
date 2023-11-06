package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// // 超时
// func groupAnagrams(strs []string) [][]string {
// 	if len(strs) == 0 {
// 		return [][]string{}
// 	}
// 	if len(strs) == 1 {
// 		return [][]string{[]string{strs[0]}}
// 	}

// 	var res = [][]string{}
// 	for len(strs) > 0 {
// 		key := strs[0]
// 		strs = strs[1:]
// 		i := 0
// 		var subRes = []string{key}
// 		for i < len(strs) {
// 			if len(strs[i]) != len(key) {
// 				i++
// 				continue
// 			}
// 			m := map[rune]int{}
// 			for _, v := range key {
// 				m[v] = m[v] + 1
// 			}
// 			var flag bool
// 			for _, v := range strs[i] {
// 				count, ok := m[v]
// 				if count <= 0 || !ok {
// 					flag = true
// 					break
// 				}
// 				m[v] = m[v] - 1
// 			}
// 			if flag {
// 				i++
// 				continue
// 			}
// 			subRes = append(subRes, strs[i])
// 			newStrs := append([]string{}, strs[0:i]...)
// 			if i+1 < len(strs) {
// 				newStrs = append(newStrs, strs[i+1:]...)
// 			}
// 			strs = newStrs
// 		}
// 		res = append(res, subRes)
// 	}
// 	return res
// }

func groupAnagrams(strs []string) [][]string {
	var res = [][]string{}
	var m = map[string][]string{}
	for _, v := range strs {
		arr := []rune(v)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		str := string(arr)
		m[str] = append(m[str], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
