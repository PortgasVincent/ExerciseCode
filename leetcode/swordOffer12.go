package leetcode

import (
	"fmt"
	"strconv"
)

func main(){
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	fmt.Println(exist1(board, "ABCCED"))
}

func exist1(board [][]byte, word string) bool {
	words := []byte(word)
	if len(words) == 0 {
		return false
	}
	if len(board) == 0 || len(board[0]) == 0{
		return false
	}
	p, q := len(board), len(board[0])
	var m = map[string]bool{}
	for i := 0; i < p; i++ {
		for j := 0; j < q; j++{
			res := dfs1(board, words, p, q, i, j, m)
			if res {
				return res
			}
		}
	}
	return false
}
func dfs1(board [][]byte, word []byte, p, q, x, y int, m map[string]bool) bool{
	if x < 0 || y < 0 ||x > p-1 || y > q-1 || m[strconv.Itoa(x) + "/" + strconv.Itoa(y)]{
		return false
	}
	if board[x][y] != word[0]{
		return false
	}
	m[strconv.Itoa(x) + "/" + strconv.Itoa(y)] = true
	word1 := word[1:]
	if len(word1) == 0 {
		return true
	}
	top := dfs1(board, word1, p, q, x, y-1, m)
	bottom := dfs1(board, word1, p, q, x, y+1, m)
	left := dfs1(board, word1, p, q, x-1, y, m)
	right := dfs1(board, word1, p, q, x+1, y, m)
	if top || bottom || left || right {
		return true
	}
	m[strconv.Itoa(x) + "/" + strconv.Itoa(y)] = false
	return false
}