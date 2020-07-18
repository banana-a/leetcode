package main

import "sort"

type tttt struct {
	key   int
	value int
}
type tttts []tttt

func (x tttts) Less(i, j int) bool {
	return x[i].value < x[j].value
}

func (x tttts) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x tttts) Len() int {
	return len(x)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*tttt, 0)
	memo := tttts{}
	for _, i2 := range nums {
		_, ok := m[i2]
		if !ok {
			m[i2] = &tttt{key: i2, value: 1}
			memo = append(memo, *m[i2])
		} else {
			m[i2].value++
		}
	}
	sort.Sort(memo)
	res := make([]int, k)
	for i := memo.Len() - 1; len(res) < k; i-- {
		res = append(res, memo[i].value)
	}
	return res
}
