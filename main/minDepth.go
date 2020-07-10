package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt32
	var help func(root *TreeNode, height int)
	help = func(root *TreeNode, height int) {
		if root.Left == nil && root.Right == nil {
			res = min(res, height)
			return
		}
		if root.Left != nil {
			help(root.Left, height+1)
		}
		if root.Right != nil {
			help(root.Right, height+1)
		}
	}
	help(root, 1)
	return res
}
