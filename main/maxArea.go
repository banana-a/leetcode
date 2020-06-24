package main

import "math"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		item := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if item > res {
			res = item
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
