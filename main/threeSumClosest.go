package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	dis := 1000000
	res := 0
	for i, num := range nums {
		left, right := i+1, len(nums)-1
		t := target - num
		for left < right {
			sum := nums[left] + nums[right]
			item := math.Abs(float64(t - sum))
			if int(item) < dis {
				dis = int(item)
				res = num + sum
			}
			if sum == t {
				return target
			}
			if sum < t {
				left++
			}
			if sum > t {
				right--
			}
		}

	}
	return res
}
