package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsWithDup(nums))
}

func restoreIpAddresses(s string) []string {
	return resplit(s, 4)
}

func resplit(s string, x int) []string {
	if len(s) == 0 {
		return nil
	}
	if x == 1 {
		if s[0] == '0' && len(s) != 1 {
			return nil
		}
		a, _ := strconv.Atoi(s)
		if a > 255 {
			return nil
		}
		return []string{s}
	}

	var res = []string{}
	if s[0] == '0' {
		sub := resplit(s[1:], x-1)
		for _, v := range sub {
			res = append(res, "0."+v)
		}
		return res
	}

	for i := 1; i <= 3 && i < len(s); i++ {
		a, _ := strconv.Atoi(s[0:i])
		if a > 255 {
			break
		}
		sub := resplit(s[i:], x-1)
		for _, v := range sub {
			res = append(res, s[0:i]+"."+v)
		}
	}
	return res
}
