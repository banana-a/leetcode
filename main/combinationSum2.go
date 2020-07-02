package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	var help func(pos int, sum int, item []int, flag bool)
	help = func(pos int, sum int, item []int, flag bool) {
		if sum == target {
			res = append(res, item)
			return
		}
		if pos >= len(candidates) || sum > target {
			return
		}
		if flag == false && pos > 0 && candidates[pos] == candidates[pos-1] {
			temp := make([]int, len(item))
			copy(temp, item)
			help(pos+1, sum, temp, false)
			return
		}
		temp1 := make([]int, len(item))
		temp2 := make([]int, len(item))
		copy(temp1, item)
		copy(temp2, item)
		temp1 = append(temp1, candidates[pos])
		help(pos+1, sum+candidates[pos], temp1, true)
		help(pos+1, sum, temp2, false)
	}
	help(0, 0, []int{}, false)
	return res
}
