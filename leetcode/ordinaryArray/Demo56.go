package ordinaryArray

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		curInterval := intervals[i]
		left := curInterval[0]
		right := curInterval[1]

		// 前一个值
		preInterval := intervals[len(intervals)-1]
		preLeft := preInterval[0]
		preRight := preInterval[1]

		if left > preRight {
			res = append(res, curInterval)
		} else {
			res[len(res)-1] = []int{preLeft, max(preRight, right)}
		}
	}
	return res
}
