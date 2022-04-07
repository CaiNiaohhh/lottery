package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	listNodeMap := make(map[*ListNode]bool, 0)
	for head != nil {
		if _, ok := listNodeMap[head]; ok {
			return head
		}
		listNodeMap[head] = true
		head = head.Next
	}
	return nil
}
