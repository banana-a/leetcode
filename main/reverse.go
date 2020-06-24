package main

import "math"

func reverse(x int) int {
	fuhao := 1
	if x < 0 {
		fuhao = -1
	}
	x = int(math.Abs(float64(x)))
	num := int64(0)
	nums := make([]int, 100)
	pos := 0
	for x > 0 {
		nums[pos] = x % 10
		x = x / 10
		pos++
	}
	temp := 0
	for temp < pos {
		num = 10*num + int64(nums[temp])
		temp++
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return int(num) * fuhao
}
