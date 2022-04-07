package main

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	queue := []*TreeNode{root}
	for true {
		length := len(queue)
		if length == 0 {
			break
		}
		for i := 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, queue[length - 1].Val)
		queue = queue[length:]
	}
	return res
}