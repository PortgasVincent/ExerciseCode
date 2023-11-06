package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	start, end := newInterval[0], newInterval[1]
	l := len(intervals)
	if l == 0 {
		return [][]int{{start, end}}
	}
	if start > intervals[l-1][1] {
		return append(intervals, newInterval)
	}
	if end < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}

	p1, p2 := find(intervals, start, 0, l-1)
	q1, q2 := find(intervals, end, 0, l-1)
	fmt.Println(p1, p2, q1, q2)

	var res [][]int
	if p1 == p2 && q1 == q2 && p1 == q1 {
		return intervals
	}
	if p1 == p2 && q1 == q2 {
		res = append(intervals[:p1], []int{intervals[p1][0], intervals[q1][1]})
		if q1+1 < l {
			res = append(res, intervals[q1+1:]...)
		}
		return res
	}
	if p1 == p2 {
		res = append(intervals[:p1], []int{intervals[p1][0], end})
		if q2 < l {
			res = append(res, intervals[q2:]...)
		}
		return res
	}
	if q1 == q2 {
		res = append(res, intervals[:p2]...)
		res := append(res, []int{start, intervals[q1][1]})
		if q1+1 < l {
			res = append(res, intervals[q1+1:]...)
		}
		return res
	}
	return append(intervals[:p2], append([][]int{{start, end}}, intervals[q2:]...)...)
}

// 如果在区间之内就返回相同index
func find(intervals [][]int, x int, left, right int) (int, int) {
	if right == left {
		if x < intervals[left][0] {
			return left - 1, left
		}
		if x > intervals[left][1] {
			return left, left + 1
		}
		return left, left
	}
	if right-left == 1 {
		if x < intervals[left][0] {
			return left - 1, left
		}
		if x <= intervals[left][1] {
			return left, left
		}
		if x < intervals[right][0] {
			return left, right
		}
		if x <= intervals[right][1] {
			return right, right
		}
		return right, right + 1
	}
	m := (left + right) / 2
	if x >= intervals[m][0] && x <= intervals[m][1] {
		return m, m
	}
	if x < intervals[m][0] {
		return find(intervals, x, left, m-1)
	}
	return find(intervals, x, m+1, right)
}
