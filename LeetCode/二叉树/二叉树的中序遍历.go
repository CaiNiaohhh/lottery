package main

// 递归解法
func _inorderTraversal(root *TreeNode) []int {
	var inorder func(root *TreeNode)
	var res []int
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res
}

// 非递归解法
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}