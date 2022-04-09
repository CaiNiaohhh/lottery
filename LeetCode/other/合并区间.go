package main

import (
	"sort"
)

// 按照左端先排序，然后逐步更新prev和cur两个区间
func _merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] > prev[1] {
			res = append(res, prev)
			prev = cur
		} else {
			if cur[1] > prev[1] {
				prev[1] = cur[1]
			}
		}
	}
	res = append(res, prev)
	return res
}
