package main

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	swap := func(a int, b int) {
		temp := nums[a]
		nums[a] = nums[b]
		nums[b] = temp
	}
	var help func(pos int)
	help = func(pos int) {
		if pos == len(nums)-1 {
			item := make([]int, len(nums))
			copy(item, nums)
			res = append(res, item)
			return
		}
		help(pos + 1)
		for i := pos + 1; i < len(nums); i++ {
			swap(pos, i)
			help(pos + 1)
			swap(i, pos)
		}
	}
	help(0)
	return res
}
