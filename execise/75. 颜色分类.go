package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	l := len(nums)
	var p, q, r int
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			p, q = i, i
			break
		}
	}
	for j := l - 1; j >= 0; j-- {
		if nums[j] != 2 {
			r = j
			break
		}
	}
	for p <= r && q <= r {
		switch nums[q] {
		case 1:
			q++
		case 0:
			if q <= p {
				q++
				continue
			}
			nums[q] = nums[p]
			nums[p] = 0
			p++
		case 2:
			nums[q] = nums[r]
			nums[r] = 2
			r--
		}
	}
	return
}

// // 2 1
//    p qr
