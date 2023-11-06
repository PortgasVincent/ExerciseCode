package main

import "strings"

func main() {

}

func romanToInt(s string) int {
	var res int
	for s != "" {
		switch {
		case strings.HasPrefix(s, "CM"):
			res += 900
			s = s[2:]
		case strings.HasPrefix(s, "CD"):
			res += 400
			s = s[2:]
		case strings.HasPrefix(s, "XC"):
			res += 90
			s = s[2:]
		case strings.HasPrefix(s, "XL"):
			res += 40
			s = s[2:]
		case strings.HasPrefix(s, "IX"):
			res += 9
			s = s[2:]
		case strings.HasPrefix(s, "IV"):
			res += 4
			s = s[2:]
		case strings.HasPrefix(s, "M"):
			res += 1000
			s = s[1:]
		case strings.HasPrefix(s, "D"):
			res += 500
			s = s[1:]
		case strings.HasPrefix(s, "C"):
			res += 100
			s = s[1:]
		case strings.HasPrefix(s, "L"):
			res += 50
			s = s[1:]
		case strings.HasPrefix(s, "X"):
			res += 10
			s = s[1:]
		case strings.HasPrefix(s, "V"):
			res += 5
			s = s[1:]
		case strings.HasPrefix(s, "I"):
			res += 1
			s = s[1:]
		}
	}

	return res
}
