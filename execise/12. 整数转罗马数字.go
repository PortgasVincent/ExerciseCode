package main

func main() {

}

func intToRoman(num int) string {
	var res string
	t := num / 1000
	for i := 0; i < t; i++ {
		res += "M"
	}
	num %= 1000
	if num >= 900 {
		res += "CM"
	}
	num %= 900
	if num >= 500 {
		res += "D"
	}
	num %= 500
	if num >= 400 {
		res += "CD"
	}
	num %= 400
	h := num / 100
	for i := 0; i < h; i++ {
		res += "C"
	}
	num %= 100
	if num >= 90 {
		res += "XC"
	}
	num %= 90
	if num >= 50 {
		res += "L"
	}
	num %= 50
	if num >= 40 {
		res += "XL"
	}
	num %= 40
	ten := num / 10
	for i := 0; i < ten; i++ {
		res += "X"
	}
	num %= 10
	if num >= 9 {
		res += "IX"
	}
	num %= 9
	if num >= 5 {
		res += "V"
	}
	num %= 5
	if num >= 4 {
		res += "IV"
	}
	num %= 4
	for i := 0; i < num; i++ {
		res += "I"
	}
	return res
}
