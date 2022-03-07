package main

/*
1. 二叉树的初始化
2. 从数组变成二叉树
	1. 前序遍历
	2. 中序遍历
	3. 后序遍历
3. 从二叉树变成数组
	1. 前序遍历
	2. 中序遍历
	3. 后序遍历
4. 层次遍历（BFS）
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 4. 层次遍历 BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	var res [][]int
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for i := 0; i < length; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, tmp)
		queue = queue[length:]
	}
	return res
}



