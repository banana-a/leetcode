package main

func isCompleteTree(root *TreeNode) bool {
	memo := make([]*TreeNode, 0)
	if root == nil {
		return true
	}
	memo = append(memo, root)
	for len(memo) != 0 {
		temp := len(memo)
		flag := false
		for i := 0; i < temp; i++ {
			item := memo[i]
			if item.Left != nil && !flag {
				memo = append(memo, item.Left)
			} else if item.Left != nil && flag {
				return false
			} else {
				flag = true
			}
			if item.Right != nil && !flag {
				memo = append(memo, item.Right)
			} else if item.Right != nil && flag {
				return false
			} else {
				flag = true
			}
		}
		memo = memo[temp:]
	}
	return true
}
