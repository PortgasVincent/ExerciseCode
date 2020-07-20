package main

func main(){

}

func longestPalindrome(s string) string {
	a := []rune(s)
	l := len(a)
	var temp string
	for i := 0; i < l; i++{
		res, length := One(a, i)
		if length > len(temp) {
			temp = res
		}
		res, length = Two(a, i)
		if length > len(temp) {
			temp = res
		}
	}
	return temp
}

func One(a []rune, i int)(res string, l int){
	res = string(a[i])
	l = 1
	for j := 0; i - j >= 0 && i + j < len(a); j++{
		if a[i-j] == a[i+j]{
			res = string(a[i-j:i+j+1])
			l = len(res)
		} else {
			return
		}
	}
	return res, l
}

func Two(a []rune, i int)(res string, l int){
	for j := 0; i - j >= 0 && i + j + 1 < len(a); j++{
		if a[i-j] == a[i+j+1]{
			res = string(a[i-j:i+j+1+1])
			l = len(res)
		} else {
			return
		}
	}
	return
}