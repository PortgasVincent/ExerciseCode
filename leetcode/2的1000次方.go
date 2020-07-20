package leetcode

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(pow2(12))
}

func pow2(n int)string{
	if n == 0 {
		return "1"
	}
	res := []int{1}
	var i = 10
	var num [][]int
	for {
		if n / i >= 1 {
			var n1 = []int{2}
			if len(num) != 0{
				n1 = num[len(num)-1]
			}
			sum := []int{1}
			for j := 0; j < 10; j++{
				sum = mul(n1, sum)
			}
			num = append(num, sum)
			i = i * 10
		} else {
			break
		}
	}
	for i <= 1 {
		if len(num) == 0{
			break
		}
		i = i / 10
		has := n / i
		for k := 0; k < has; k++{
			res = mul(res, num[len(num)-1])
		}
		n = n % i
		num = num[:len(num)-1]
	}
	for m := 0; m < n; m++{
		res = mul(res, []int{2})
	}
	return getRes(res)
}

func getRes(res []int) string {
	l := len(res)
	resStr := ""
	for i := l-1; i >= 0; i--{
		resStr += strconv.Itoa(res[i])
	}
	return resStr
}

//数字倒序

func mul(p1 []int, p2 []int) []int{
	res := []int{}
	temp :=  0
	pos := 0
	for i := 0; i < len(p1); i++{
		for j := 0; j < len(p2); j++{
			r := p1[i] * p2[j]
			r += temp
			temp = r / 10
			pos = i + j
			if len(res)-1 < pos{
				res = append(res, r%10)
			} else {
				sum := res[pos] + (r%10)
				res[pos] = sum%10
				temp += sum/10
			}

			for temp != 0{
				if len(res)-1 < pos+1{
					res = append(res, temp%10)
					temp = temp / 10
				} else {
					sum := res[pos+1] + temp%10
					res[pos+1] = sum%10
					temp = temp / 10
					temp += sum/10
				}
			}
		}

	}
	return res
}