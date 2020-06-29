package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(nums[left:right+1], nums[mid])
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target < nums[0] {
				left = mid + 1
			} else if target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if nums[0] > nums[mid] {
			if target >= nums[0] {
				right = mid - 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return -1
}
