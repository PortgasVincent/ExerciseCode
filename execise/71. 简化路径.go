package main

import (
	"fmt"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	l := len(path)
	new := "/"
	pre := []int{}
	cur := 0
	i := 1
	for ; i < l; i++ {
		if string(path[i]) == "/" {
			if i-cur == 1 {
				cur = i
				continue
			}
			temp := path[cur+1 : i]
			switch temp {
			case ".":
			case "..":
				lPre := len(pre)
				if lPre != 0 {
					new = new[:pre[lPre-1]]
					pre = pre[:lPre-1]
				}
			default:
				pre = append(pre, len(new))
				new += temp + "/"
			}
			cur = i
		}
	}
	temp := path[cur+1 : i]
	switch temp {
	case ".":
	case "..":
		lPre := len(pre)
		if lPre != 0 {
			new = new[:pre[lPre-1]]
			pre = pre[:lPre-1]
		}
	default:
		pre = append(pre, len(new))
		new += temp
	}
	if len(new) > 1 && string(new[len(new)-1]) == "/" {
		new = new[:len(new)-1]
	}
	return new
}
