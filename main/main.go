package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var nums []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	var target int
	fmt.Scan(&target)
	fmt.Println(twoSum(nums, target))
}
