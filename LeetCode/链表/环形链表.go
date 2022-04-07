package main

// 直接搞个哈希表就行
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	Map := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := Map[head]; ok {
			return true
		}
		Map[head] = true
		head = head.Next
	}
	return false
}

// 快慢指针
func _hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}