package main

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	x := len(text1)
	y := len(text2)
	m := make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		m[i] = make([]int, y+1)
	}

	for i := 1; i < x+1; i++ {
		for j := 1; j < y+1; j++ {
			if text1[i-1] == text2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				m[i][j] = max(max(m[i-1][j], m[i][j-1]), m[i-1][j-1])
			}
		}
	}
	return m[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
