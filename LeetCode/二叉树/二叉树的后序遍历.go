package main

// 递归
func _postorderTraversal(root *TreeNode) []int {
	var postOrder func(root *TreeNode)
	var res []int
	postOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrder(root.Left)
		postOrder(root.Right)
		res = append(res, root.Val)
	}
	postOrder(root)
	return res
}

// 非递归
// 前序遍历的相反 根右左 之后倒排
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack) - 1].Left
		stack = stack[:len(stack) - 1]
	}
	reverse := func(Slice []int) []int {
		for i, j := 0, len(Slice) - 1; i < j; i, j = i + 1, j - 1 {
			Slice[i], Slice[j] = Slice[j], Slice[i]
		}
		return Slice
	}
	return reverse(res)
}