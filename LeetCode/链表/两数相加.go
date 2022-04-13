package main
/*
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，
并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	head := node
	le := 0
	for l1 != nil || l2 != nil {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		res := l1Val + l2Val + le
		newNode := new(ListNode)
		newNode.Val = res % 10
		le = res / 10
		node.Next = newNode
		node = node.Next
	}
	if le != 0 {
		node.Next = &ListNode{Val: le}

	}
	return head.Next
}
