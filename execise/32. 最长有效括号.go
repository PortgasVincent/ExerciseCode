package main

func main() {

}

func longestValidParentheses(s string) int {
	l := len(s)
	var max int = 0
	for i := 0; i < l; i++ {
		temp := 0
		tempRes := 0
		leftCount := 0
	f:
		for j := i; j < l; j++ {
			switch string(s[j]) {
			case "(":
				leftCount++
				temp++
			case ")":
				if leftCount == 0 {
					break f
				}
				leftCount--
				temp++
			}
			if leftCount == 0 {
				tempRes = temp
			}
		}
		if tempRes > max {
			max = tempRes
		}
	}
	return max
}
