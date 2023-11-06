package main

func main() {

}

func convert(s string, numRows int) string {
	arr := []rune(s)
	if len(arr) <= 1 || numRows <= 1 {
		return s
	}
	strArr := make([]string, numRows)
	flag := 1
	j := 0
	for i := 0; i < len(arr); i++ {
		strArr[j] += string(arr[i])
		j += flag
		if j == numRows-1 || j == 0 {
			flag = -flag
		}
	}
	var res string
	for _, v := range strArr {
		res += v
	}
	return res
}
