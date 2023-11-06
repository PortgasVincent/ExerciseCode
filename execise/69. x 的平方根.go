package main

import (
	"fmt"
)

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	for i := 1; i <= x/2; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			return i
		}
	}
	return x
}
