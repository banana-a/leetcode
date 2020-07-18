package main

func rightSideView(root *TreeNode) []int {
	memo := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return nil
	}
	memo = append(memo, root)
	for len(memo) != 0 {
		num := len(memo)
		res = append(res, memo[0].Val)
		for i := 0; i < num; i++ {
			item := memo[0]
			if item.Right != nil {
				memo = append(memo, item.Right)
			}
			if item.Left != nil {
				memo = append(memo, item.Left)
			}
			memo = memo[1:]
		}
	}
	return res
}
