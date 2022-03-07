package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	var queue []*TreeNode
	dir := false
	queue = append(queue, root)
	for len(queue) > 0 {
		var tmp []int
		length := len(queue)
		for i := 0; i < length; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if dir {
			_ = rev(tmp)
			dir = false
		} else {
			dir = true
		}
		ans = append(ans, tmp)
		queue = queue[length:]
	}
	return ans
}

// 反转切片
func rev(Slice []int) []int {
	for i, j := 0, len(Slice) - 1; i < j; i, j = i + 1, j - 1 {
		Slice[i], Slice[j] = Slice[j], Slice[i]
	}
	return Slice
}