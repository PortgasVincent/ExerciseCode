package main

import "math"

func main() {

}

func reverse(x int) int {
	var res = 0
	for x != 0 {
		res = res*10 + x%10
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}
