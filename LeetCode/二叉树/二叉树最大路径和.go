package main


func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max = -1 << 31
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := Max(dfs(node.Left), 0)
		rightMax := Max(dfs(node.Right), 0)
		max = Max(max, node.Val + leftMax + rightMax)
		return node.Val + Max(leftMax, rightMax)
	}
	dfs(root)
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}