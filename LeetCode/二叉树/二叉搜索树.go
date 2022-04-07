package main

/*给定二叉搜索树（BST）的根节点root和一个整数值val。
你需要在 BST 中找到节点值等于val的节点。 返回以该节点为根的子树。
如果节点不存在，则返回null。
*/


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
