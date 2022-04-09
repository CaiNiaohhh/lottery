package main
// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := &ListNode{0, head}
	pre, cur := node, head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			val := cur.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return pre.Next
}