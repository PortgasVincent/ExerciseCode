package main

func main() {

}

// func isMatch(s string, p string) bool {
// 	if s == "" && p == "" {
// 		return true
// 	}
// 	if s == "" || p == "" {
// 		return false
// 	}

// 	// prepare
// 	arrS := []rune(s)
// 	arrP := []rune(p)
// 	ls := len(arrS)
// 	lp := len(arrP)
// 	m := [][]bool{}
// 	for i := 0; i < ls+1; i++ {
// 		m = append(m, []bool{})
// 		for j := 0; j < lp+1; j++ {
// 			m[i] = append(m[i], false)
// 		}
// 	}

// 	// base case
// 	m[0][0] = true
// 	for j := 1; j < lp+1; j++ {
// 		if string(arrP[j-1]) == "*" {
// 			m[0][j] = m[0][j-2]
// 		}
// 	}

// 	// 迭代
// 	for i := 1; i < ls+1; i++ {
// 		for j := 1; j < lp+1; j++ {
// 			if arrS[i-1] == arrP[j-1] || string(arrP[j-1]) == "." {
// 				m[i][j] = m[i-1][j-1]
// 			} else if string(arrP[j-1]) == "*" {
// 				if arrS[i-1] == arrP[j-2] || string(arrP[j-2]) == "." {
// 					m[i][j] = m[i-1][j] || m[i-1][j-2] || m[i][j-2]
// 					// p的最后两位使用次数：
// 					// 0: m[i][j-2], 1: m[i-1][j-2], 多次: m[i-1][j]
// 				} else {
// 					m[i][j] = m[i][j-2]
// 				}
// 			}
// 		}
// 	}
// 	return m[ls][lp]
// }

func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)

	m := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		m[i] = make([]bool, lp+1)
	}
	m[0][0] = true
	for i := 1; i <= lp; i++ {
		if string(p[i-1]) == "*" {
			m[0][i] = m[0][i-2]
		}
	}

	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if s[i-1] == p[j-1] || string(p[j-1]) == "." {
				m[i][j] = m[i-1][j-1]
			} else if string(p[j-1]) == "*" {
				if string(p[j-2]) == "." || p[j-2] == s[i-1] {
					m[i][j] = m[i][j-2] || m[i-1][j-2] || m[i-1][j]
				} else {
					m[i][j] = m[i][j-2]
				}
			}
		}
	}
	return m[ls][lp]
}
