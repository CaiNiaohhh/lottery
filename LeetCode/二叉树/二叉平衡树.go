package main

import "math"

// 实现一个函数，检查二叉树是否平衡，平衡的定义为：任意一个节点，其两颗子树的高度差不超过1

// 首先得定义一个深度的函数，计算每个节点的深度为多少

func depth(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(depth(root.Left), depth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Left) &&
		isBalanced(root.Right) &&
		math.Abs(depth(root.Left) - depth(root.Right)) <= 1
}
