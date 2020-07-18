package main

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var helpPathSum func(root *TreeNode, now int, item []int)
	helpPathSum = func(root *TreeNode, now int, item []int) {
		if root.Left == nil && root.Right == nil {
			if sum == now+root.Val {
				item = append(item, root.Val)
				var temp = make([]int, len(item))
				copy(temp, item)
				res = append(res, temp)
				item = item[0 : len(item)-1]
			} else {
				return
			}
		}
		if root.Left != nil {
			item = append(item, root.Val)
			helpPathSum(root.Left, now+root.Val, item)
			item = item[0 : len(item)-1]
		}
		if root.Right != nil {
			item = append(item, root.Val)
			helpPathSum(root.Right, now+root.Val, item)
			item = item[0 : len(item)-1]
		}
	}
	helpPathSum(root, 0, []int{})
	return res
}
