package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m := map[string]bool{}
	return match(s1, s2, s3, m)
}

func match(s1, s2, s3 string, m map[string]bool) bool {
	if v, ok := m[s1+","+s2+","+s3]; ok {
		return v
	}
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if l1 == 0 && l2 == 0 && l3 == 0 {
		return true
	}
	if l3 != l1+l2 {
		return false
	}
	i, j, k := 0, 0, 0
	for k < l3 {
		if i < l1 && s3[k] == s1[i] && j < l2 && s3[k] == s2[j] {
			res := match(s1[i+1:], s2[j:], s3[k+1:], m) || match(s1[i:], s2[j+1:], s3[k+1:], m)
			m[s1+","+s2+","+s3] = res
			return res
		}
		if i < l1 && s3[k] == s1[i] {
			i++
		} else if j < l2 && s3[k] == s2[j] {
			j++
		} else {
			break
		}
		k++
	}
	if k < l3 || i < l1 || j < l2 {
		m[s1+","+s2+","+s3] = false
		return false
	}
	m[s1+","+s2+","+s3] = true
	return true
}
