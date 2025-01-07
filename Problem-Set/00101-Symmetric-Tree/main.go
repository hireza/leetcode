package main

// TreeNode represents a node in a tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}

	return isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
