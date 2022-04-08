package main

/*
	给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
	请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/

// 想法是找到left和right对应的两个节点 但是在题意里面对应的是left - 2 和right

// 定义反转链表的函数，但是返回head和tail
func reverseList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ans := &ListNode{Val: -1}
	ans.Next = head
	pre := ans
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}
	rightNode := pre
	for i := 0; i < right - left + 1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	curNode := rightNode.Next

	// 切割
	pre.Next = nil
	rightNode.Next = nil

	reverseList(leftNode)

	// 重新接上
	pre.Next = rightNode
	leftNode.Next = curNode
	return ans.Next
}