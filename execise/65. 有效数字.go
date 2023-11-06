package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

var numMap = map[string]bool{
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

var prefix = map[string]bool{
	"-": true,
	"+": true,
}

var ePrefix = map[string]bool{
	"e": true,
	"E": true,
}

func isNumber(s string) bool {
	var index = -1
	for i, v := range s {
		if ePrefix[string(v)] {
			if i == len(s)-1 {
				return false
			}
			index = i
			break
		}
	}
	if index == -1 {
		return isInt(s) || isFloat(s)
	}
	if index == 0 {
		return false
	}
	return (isInt(s[:index]) || isFloat(s[:index])) && isInt(s[index+1:])
}

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	if prefix[string(s[0])] {
		s = s[1:]
	}
	countNum := 0
	for _, v := range s {
		if numMap[string(v)] {
			countNum++
			continue
		}
		return false
	}
	if countNum == 0 {
		return false
	}
	return true
}

func isFloat(s string) bool {
	if len(s) == 0 {
		return false
	}
	if prefix[string(s[0])] {
		s = s[1:]
	}
	countNum := 0
	countDot := 0
	for _, v := range s {
		if string(v) == "." {
			if countDot == 0 {
				countDot++
				continue
			}
			return false
		}
		if numMap[string(v)] {
			countNum++
			continue
		}
		return false
	}
	if countNum > 0 {
		return true
	}
	return false
}
