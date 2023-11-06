package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func addBinary(a string, b string) string {
	l1 := len(a)
	l2 := len(b)

	var res []byte
	var flag int
	var i, j = l1 - 1, l2 - 1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		temp1, temp2 := int(a[i]-'0'), int(b[j]-'0')
		sum := temp1 + temp2 + flag
		flag = sum / 2
		res = append([]byte{byte(sum%2 + '0')}, res...)
	}
	for ; i >= 0; i-- {
		sum := int(a[i]-'0') + flag
		flag = sum / 2
		res = append([]byte{byte(sum%2 + '0')}, res...)
	}
	for ; j >= 0; j-- {
		sum := int(b[j]-'0') + flag
		flag = sum / 2
		res = append([]byte{byte(sum%2 + '0')}, res...)
	}
	if flag == 1 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}
