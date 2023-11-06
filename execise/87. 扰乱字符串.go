package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isScramble("great", "rgeat"))
}

func isScramble(s1 string, s2 string) bool {
	l := len(s1)
	m := map[string]bool{}
	return isMatch(s1, s2, 0, l-1, 0, l-1, m)
}

func isMatch(s1, s2 string, left1, right1, left2, right2 int, m map[string]bool) bool {
	v, ok := m[getKey(left1, right1, left2, right2)]
	if ok {
		return v
	}

	l := right1 - left1 + 1
	if l == 1 {
		if s1[left1:right1+1] == s2[left2:right2+1] {
			m[getKey(left1, right1, left2, right2)] = true
			return true
		}
		m[getKey(left1, right1, left2, right2)] = false
		return false
	}

	for i := 1; i < l; i++ {
		m1 := map[string]int{}
		for j := left1; j < left1+i; j++ {
			m1[string(s1[j])] = m1[string(s1[j])] + 1
		}
		m2 := map[string]int{}
		for j := left1 + i; j <= right1; j++ {
			m2[string(s1[j])] = m2[string(s1[j])] + 1
		}
		if compare(s2[left2:left2+i], m1) && compare(s2[left2+i:right2+1], m2) && isMatch(s1, s2, left1, left1+i-1, left2, left2+i-1, m) && isMatch(s1, s2, left1+i, right1, left2+i, right2, m) {
			m[getKey(left1, right1, left2, right2)] = true
			return true
		}
		if compare(s2[right2-i+1:right2+1], m1) && compare(s2[left2:right2-i+1], m2) && isMatch(s1, s2, left1, left1+i-1, right2-i+1, right2, m) && isMatch(s1, s2, left1+i, right1, left2, right2-i, m) {
			m[getKey(left1, right1, left2, right2)] = true
			return true
		}
	}
	m[getKey(left1, right1, left2, right2)] = false
	return false
}

func getKey(left1, right1, left2, right2 int) string {
	return strconv.Itoa(left1) + "," + strconv.Itoa(right1) + "," + strconv.Itoa(left2) + "," + strconv.Itoa(right2)
}

func compare(s string, m map[string]int) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		v, ok := m[string(s[i])]
		if !ok || v <= 0 {
			return false
		}
	}
	return true
}
