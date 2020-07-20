package leetcode

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(maxProbability(5, [][]int{
		{1,4},
		{2,4},
		{0,4},
		{0,3},
		{0,2},
		{2,3},
	}, []float64{0.37,0.17,0.93,0.23,0.39,0.04}, 3, 4))
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	set := map[string]float64{}
	for i := 0; i < len(edges); i++{
		set[getKey(edges[i][0], edges[i][1])] = succProb[i]
	}
	var res float64 = set[getKey(start, end)]
	nodes := []int{}
	for i := 0; i < n; i++{
		nodes = append(nodes, i)
	}
	getProbability(nodes, set, &res, 1, start, end)
	return res
}

func getProbability(nodes []int, set map[string]float64, res *float64, current float64, start, end int){
	fmt.Printf("start: %v, res: %v, current: %v\n", start, *res, current)
	if current <= *res {
		return
	}
	if start == end && current > *res{
		*res = current
	}
	newNodes := []int{}
	for _, v := range nodes {
		if v != start {
			newNodes = append(newNodes, v)
		}
	}
	for _, v := range newNodes {
		getProbability(newNodes, set, res, FloatRound(set[getKey(start, v)] * current, 5), v, end)
	}
	return
}

func getKey(p1, p2 int)string{
	if p1 <= p2{
		return strconv.Itoa(p1) + "/" + strconv.Itoa(p2)
	}
	return strconv.Itoa(p2) + "/" + strconv.Itoa(p1)
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}