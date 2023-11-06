package main

func main() {

}

func solveSudoku(board [][]byte) {
	var line, col [9][9]bool
	var block [3][3][9]bool
	var space [][2]int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
			} else {
				digit := board[i][j] - '1'
				line[i][digit] = true
				col[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == len(space) {
			return true
		}
		i, j := space[pos][0], space[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !col[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				col[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				col[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
	return
}
