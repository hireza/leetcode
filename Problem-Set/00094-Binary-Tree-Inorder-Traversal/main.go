package main

// TreeNode represents a node in a tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// // RECURSIVE WAY
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	inorderHelper(root, &res)
// 	return res
// }

// func inorderHelper(node *TreeNode, res *[]int) {
// 	if node == nil {
// 		return
// 	}

// 	inorderHelper(node.Left, res)
// 	*res = append(*res, node.Val)
// 	inorderHelper(node.Right, res)
// }

// ITERATIVE WAY
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		current = current.Right
	}

	return res
}
