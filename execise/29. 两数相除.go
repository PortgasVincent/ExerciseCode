package main

import (
	"math"
)

func main() {

}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	var negative = false
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		negative = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < divisor {
		return 0
	}

	var mul, rest = 1, divisor
	for rest <= dividend-rest && rest <= math.MaxInt32-rest {
		rest += rest
		mul += mul
	}

	res := divide(dividend-rest, divisor)
	if negative {
		if math.MinInt32+res > -mul {
			return math.MinInt32
		}
		return -(res + mul)
	}

	if math.MaxInt32-res < mul {
		return math.MaxInt32
	}
	return mul + res
}
