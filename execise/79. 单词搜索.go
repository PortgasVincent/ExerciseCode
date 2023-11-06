package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	checkM := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				checkM[getKey(i, j)] = true
				res := find(board, word[1:], i, j, checkM)
				if res {
					return true
				}
				checkM[getKey(i, j)] = false
			}
		}
	}
	return false
}

func find(board [][]byte, word string, x, y int, m map[string]bool) bool {
	if len(word) == 0 {
		return true
	}
	lx := len(board)
	ly := len(board[0])
	var res bool
	if x-1 >= 0 && !m[getKey(x-1, y)] && board[x-1][y] == word[0] {
		m[getKey(x-1, y)] = true
		res = find(board, word[1:], x-1, y, m)
		if res {
			return true
		}
		m[getKey(x-1, y)] = false
	}
	if y-1 >= 0 && !m[getKey(x, y-1)] && board[x][y-1] == word[0] {
		m[getKey(x, y-1)] = true
		res = find(board, word[1:], x, y-1, m)
		if res {
			return true
		}
		m[getKey(x, y-1)] = false
	}
	if x+1 < lx && !m[getKey(x+1, y)] && board[x+1][y] == word[0] {
		m[getKey(x+1, y)] = true
		res = find(board, word[1:], x+1, y, m)
		if res {
			return true
		}
		m[getKey(x+1, y)] = false
	}
	if y+1 < ly && !m[getKey(x, y+1)] && board[x][y+1] == word[0] {
		m[getKey(x, y+1)] = true
		res = find(board, word[1:], x, y+1, m)
		if res {
			return true
		}
		m[getKey(x, y+1)] = false
	}
	return false
}

func getKey(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
