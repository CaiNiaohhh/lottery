package main

// 递归
func _preorderTraversal(root *TreeNode) []int {
	var preOrder func(root *TreeNode)
	var res []int
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return res
}

// 非递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1].Right
		stack = stack[:len(stack) - 1]
	}
	return res
}