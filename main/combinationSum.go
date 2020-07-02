package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var help func(pos int, sum int, nums []int)
	help = func(pos int, sum int, nums []int) {
		if sum == target {
			res = append(res, nums)
			return
		}
		if sum > target || pos >= len(candidates) {
			return
		}
		if sum < target {
			for i := 0; i*candidates[pos]+sum <= target; i++ {
				item := []int{}
				for _, num := range nums {
					item = append(item, num)
				}
				for t := 0; t < i; t++ {
					item = append(item, candidates[pos])
				}
				help(pos+1, sum+i*candidates[pos], item)
			}
		}
	}
	for i := 0; i < len(candidates); i++ {
		for j := 1; j*candidates[i] <= target; j++ {
			item := []int{}
			for t := 0; t < j; t++ {
				item = append(item, candidates[i])
			}
			help(i+1, j*candidates[i], item)
		}
	}
	return res
}
