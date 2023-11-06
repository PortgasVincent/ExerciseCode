package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

var rela = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

func numDecodings(s string) int {
	m := map[string]int{}
	return deco(s, m)
}

func deco(s string, m map[string]int) int {
	if v, ok := m[s]; ok {
		return v
	}

	l := len(s)
	if l == 0 || s[0] == '0' {
		return 0
	}
	if l == 1 {
		return 1
	}
	if l == 2 {
		var res = 0
		if rela[s] {
			res++
		}
		if s[1] != '0' {
			res++
		}
		return res
	}

	var sum int
	if s[1] != '0' {
		sum = sum + deco(s[1:], m)
	}
	if l > 2 && s[2] != '0' && rela[s[0:2]] {
		sum += deco(s[2:], m)
	}
	m[s] = sum
	return sum
}
