package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow(2, 10))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	var negative bool
	if x < 0 && n%2 != 0 {
		negative = true
	}
	var flag bool
	if n < 0 {
		flag = true
	}

	x = math.Abs(x)
	n = int(math.Abs(float64(n)))

	var m = map[int]float64{
		1: x,
	}

	var temp = x
	i := 2
	for i < n {
		temp = temp * temp
		m[i] = temp
		i *= 2
	}
	i /= 2

	var sum float64 = 1
	for n > 0 {
		if n >= i {
			sum *= m[i]
			n = n - i
		} else {
			i = i / 2
		}
	}

	if flag {
		sum = 1 / sum
	}
	if negative {
		sum = -sum
	}
	return sum
}
