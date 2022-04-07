package main

// 可以使用分治法，也就是归并排序的思想
// 合并两个有序链表
func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	tmp := &ListNode{}
	head := tmp
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil && l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else if l1 != nil && l2 != nil && l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else if l1 != nil {
			tmp.Next = l1
			break
		} else if l2 != nil {
			tmp.Next = l2
			break
		}
		tmp = tmp.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var Sort func(left, right int) *ListNode
	Sort = func(left, right int) *ListNode {
		if left == right {
			return lists[left]
		}
		mid := left + (right - left) >> 1
		l := Sort(left, mid)
		r := Sort(mid + 1, right)
		return mergeTwoList(l, r)
	}
	return Sort(0, len(lists) - 1)
}
