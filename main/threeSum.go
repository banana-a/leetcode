package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var help func(int, int, int)
	help = func(pos int, i int, j int) {
		target := -nums[pos]
		if i >= j {
			return
		}
		if nums[i] == nums[i-1] && i != pos+1 {
			help(pos, i+1, j)
			return
		}
		if j != len(nums)-1 && nums[j] == nums[j+1] {
			help(pos, i, j-1)
			return
		}
		sum := nums[i] + nums[j]
		if sum == target {
			item := make([]int, 3)
			item[0] = nums[pos]
			item[1] = nums[i]
			item[2] = nums[j]
			res = append(res, item)
			help(pos, i+1, j-1)
		} else if sum < target {
			help(pos, i+1, j)
		} else if sum > target {
			help(pos, i, j-1)
		}
	}
	for i, _ := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		help(i, i+1, len(nums)-1)
	}
	return res
}
