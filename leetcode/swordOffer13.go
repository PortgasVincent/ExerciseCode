package leetcode

import (
	"fmt"
	"strconv"
)

var (
	ans = 0
	exist = map[string]bool{}
)
func main(){
	dfs(3,2,17, 0, 0)
	fmt.Println(ans)
}

func dfs(m, n, k, x, y int){
	if x < 0 || x >= m || y < 0 || y >= n || exist[strconv.Itoa(x)+"/"+strconv.Itoa(y)]{
		return
	}
	temp := 0
	x1, y1 := x, y
	for x1 > 0{
		temp += x1%10
		x1 /= 10
	}
	for y1 > 0{
		temp += y1%10
		y1 /= 10
	}
	if temp > k{
		return
	}
	dfs(m,n,k,x+1,y)
	dfs(m,n,k,x,y+1)
	exist[strconv.Itoa(x)+"/"+strconv.Itoa(y)] = true
	ans ++
	return
}
