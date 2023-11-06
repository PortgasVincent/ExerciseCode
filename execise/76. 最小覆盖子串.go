package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func minWindow(s string, t string) string {
	m := map[string]int{}
	for _, v := range t {
		m[string(v)] = m[string(v)] + 1
	}
	count := 0
	l := len(t)
	min := math.MaxInt32
	res := ""
	for i, j := 0, 0; j < len(s); j++ {
		if v, ok := m[string(s[j])]; ok {
			m[string(s[j])] = v - 1
			if v <= 0 {
				continue
			}
			count++
			if count == l {
				for v1, ok1 := m[string(s[i])]; !ok1 || v1 < 0; {
					if ok1 {
						m[string(s[i])] = v1 + 1
					}
					i++
					v1, ok1 = m[string(s[i])]
				}
				if j-i+1 < min {
					res = s[i : j+1]
					min = j - i + 1
				}
				m[string(s[i])] = m[string(s[i])] + 1
				count--
				i++
			}
		}
	}
	return res
}
