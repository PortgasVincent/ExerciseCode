package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	m := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var res []string
	for i := 0; i < len(digits); i++ {
		str := m[string(digits[i])]
		var temp = []string{}
		for _, v := range str {
			if len(res) == 0 {
				temp = append(temp, string(v))
				continue
			}
			for _, s := range res {
				temp = append(temp, s+string(v))
			}
		}
		res = temp
	}
	return res
}
