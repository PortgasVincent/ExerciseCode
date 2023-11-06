package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var mx = map[byte]bool{}
		var my = map[byte]bool{}
		for j := 0; j < 9; j++ {
			if (mx[board[i][j]] && board[i][j] != '.') || (my[board[j][i]] && board[j][i] != '.') {
				return false
			}
			mx[board[i][j]] = true
			my[board[j][i]] = true
		}
	}
	for i := 0; i+3-1 < 9; i += 3 {
		for j := 0; j+3-1 < 9; j += 3 {
			var m = map[byte]bool{}
			for p := i; p < i+3; p++ {
				for q := j; q < j+3; q++ {
					if m[board[p][q]] && board[p][q] != '.' {
						return false
					}
					m[board[p][q]] = true
				}
			}
		}
	}
	return true
}
