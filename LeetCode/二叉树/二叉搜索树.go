package main

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	for root != nil {
		if root.Val == val {
			return root
		}
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}
