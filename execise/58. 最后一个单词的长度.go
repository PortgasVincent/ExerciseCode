package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	l := len(s)
	var count = 0
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		}
		count++
	}
	return count
}
