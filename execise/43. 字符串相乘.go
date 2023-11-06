package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var sum string
	for i := 0; i < len(num1); i++ {
		sum = add(sum+"0", singleMul(num1[i], num2))
	}

	return sum
}

func singleMul(single byte, num string) string {
	if single == '0' {
		return "0"
	}
	var res string
	var large10count int
	for i := len(num) - 1; i >= 0; i-- {
		temp := int(single-'0')*int(num[i]-'0') + large10count
		large10count = temp / 10
		res = string(temp%10+'0') + res
	}
	if large10count > 0 {
		res = string(large10count+'0') + res
	}
	return res
}

func add(num1 string, num2 string) string {
	var res string
	l1 := len(num1)
	l2 := len(num2)
	i, j := l1-1, l2-1

	var large10count int
	for i >= 0 || j >= 0 || large10count > 0 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		temp := a + b + large10count
		res = string(temp%10+'0') + res
		large10count = temp / 10
		i--
		j--
	}

	return res
}
