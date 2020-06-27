package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i, one := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			two := nums[j]
			if two == nums[j-1] && j != i+1 {
				continue
			}
			start := j + 1
			end := len(nums) - 1
			t := target - one - two
			for start < end {
				sum := nums[start] + nums[end]
				if sum == t {
					res = append(res, []int{one, two, nums[start], nums[end]})
					start++
					end--
					for nums[start] == nums[start]-1 && start != j+1 && start < end {
						start++
					}
					for end != len(nums)-1 && nums[end] == nums[end+1] && start < end {
						end--
					}
				} else if sum < t {
					start++
				} else if sum > t {
					end--
				}
			}
		}
	}
	return res
}
