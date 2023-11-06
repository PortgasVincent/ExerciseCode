package main

import (
	"fmt"
)

func main() {
	fmt.Println(getPermutation(3, 1))
}

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	mulMap := make([]int, n)
	mulMap[0] = 1
	for i := 2; i <= n; i++ {
		mulMap[i-1] = mulMap[i-2] * i
	}
	m := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	var res string
	for i := n - 2; i >= 0; i-- {
		temp := k / mulMap[i]
		k = k % mulMap[i]

		if k == 0 {
			res += m[arr[temp-1]]
			if temp == len(arr) {
				arr = arr[:temp-1]
			} else {
				arr = append(arr[:temp-1], arr[temp:]...)
			}
			for j := len(arr) - 1; j >= 0; j-- {
				res += m[arr[j]]
			}
			return res
		}

		temp += 1
		res += m[arr[temp-1]]
		if temp == len(arr) {
			arr = arr[:temp-1]
		} else {
			arr = append(arr[:temp-1], arr[temp:]...)
		}
	}
	return res
}
