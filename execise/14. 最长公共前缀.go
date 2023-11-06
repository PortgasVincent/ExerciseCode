package main

import "strings"

func main() {

}

func longestCommonPrefix(strs []string) string {
	var common string = strs[0]
	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], common) && common != "" {
			common = common[:len(common)-1]
		}
	}
	return common
}
