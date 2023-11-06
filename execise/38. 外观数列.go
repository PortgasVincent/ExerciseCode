package main

import "fmt"

func main() {
	fmt.Println(countAndSay(2))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var res = "1"
	for i := 1; i < n; i++ {
		res = count(res)
	}
	return res
}

func count(s string) string {
	var res, last string
	var count int
	for _, v := range s {
		if last == "" {
			last = string(v)
			count++
			continue
		}
		if string(v) != last {
			res = res + string(count+'0') + last
			last = string(v)
			count = 1
			continue
		}
		count++
	}
	res = res + string(count+'0') + last
	return res
}
