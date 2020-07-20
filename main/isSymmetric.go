package main

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelp(root, root)
}

func isSymmetricHelp(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricHelp(p.Left, q.Right) && isSymmetricHelp(p.Right, q.Left)
}
