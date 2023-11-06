package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{[]}"))
}

func isValid(s string) bool {
	left := map[string]bool{
		"(": true,
		"{": true,
		"[": true,
	}
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var arr []string
	for _, v := range s {
		if left[string(v)] {
			arr = append(arr, string(v))
			continue
		}
		l := len(arr)
		if l == 0 {
			return false
		}

		if arr[l-1] != m[string(v)] {
			return false
		}
		arr = arr[:l-1]
	}
	if len(arr) != 0 {
		return false
	}
	return true
}
