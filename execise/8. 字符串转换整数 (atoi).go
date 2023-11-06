package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.MinInt32)
}

func myAtoi(s string) int {
	for strings.HasPrefix(s, " ") {
		s = strings.TrimPrefix(s, " ")
	}

	arr := []rune(s)
	negative := false

	if len(arr) > 0 && (string(arr[0]) == "-" || string(arr[0]) == "+") {
		switch string(arr[0]) {
		case "+":
			negative = false
		case "-":
			negative = true
		}
		arr = arr[1:]
	}

	var res = 0
	var exceed bool
f1:
	for _, v := range arr {
		var subExceed, equal bool
		if res > math.MaxInt32/10 {
			subExceed = true
		}
		if res == math.MaxInt32/10 {
			equal = true
		}

		switch string(v) {
		case "0":
			res = res * 10
			if subExceed {
				exceed = true
			}
		case "1":
			res = res*10 + 1
			if subExceed {
				exceed = true
			}
		case "2":
			res = res*10 + 2
			if subExceed {
				exceed = true
			}
		case "3":
			res = res*10 + 3
			if subExceed {
				exceed = true
			}
		case "4":
			res = res*10 + 4
			if subExceed {
				exceed = true
			}
		case "5":
			res = res*10 + 5
			if subExceed {
				exceed = true
			}
		case "6":
			res = res*10 + 6
			if subExceed {
				exceed = true
			}
		case "7":
			res = res*10 + 7
			if subExceed {
				exceed = true
			}
		case "8":
			res = res*10 + 8
			if subExceed || (equal && !negative) {
				exceed = true
			}
		case "9":
			res = res*10 + 9
			if subExceed || equal {
				exceed = true
			}
		default:
			break f1
		}
	}

	if exceed {
		if negative {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if negative {
		res = -res
	}

	return res
}
