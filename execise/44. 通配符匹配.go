package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("", "****"))
}

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	if p == "" {
		return false
	}

	ls := len(s)
	lp := len(p)
	m := [][]bool{}
	for i := 0; i < ls+1; i++ {
		m = append(m, []bool{})
		for j := 0; j < lp+1; j++ {
			m[i] = append(m[i], false)
		}
	}

	m[0][0] = true
	for j := 1; j <= len(p) && string(p[j-1]) == "*"; j++ {
		m[0][j] = true
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if s[i-1] == p[j-1] || string(p[j-1]) == "?" {
				m[i][j] = m[i-1][j-1]
			} else if string(p[j-1]) == "*" {
				m[i][j] = m[i-1][j] || m[i][j-1]
			}
		}
	}

	return m[ls][lp]
}
