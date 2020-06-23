package main
func twoSum(nums []int, target int) []int {
	res := make([]int,2)
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i;
	}
	for i, num := range nums {
		var index,ok = m[target - num]
		if ok && index != i{
			res[0] = i;
			res[1] = m[target - num]
			break
		}
	}
	return res
}