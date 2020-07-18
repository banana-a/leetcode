package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	memo := make([]*TreeNode, 0)
	memo = append(memo, root)
	num := 1
	for len(memo) != 0 {
		num = len(memo)
		item := make([]int, 0)
		for i := 0; i < num; i++ {
			temp := memo[0]
			item = append(item, temp.Val)
			memo = memo[1:]
			if temp.Left != nil {
				memo = append(memo, temp.Left)
			}
			if temp.Right != nil {
				memo = append(memo, temp.Right)
			}
		}
		res = append(res, item)
	}
	return res
}
