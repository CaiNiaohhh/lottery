package main

import "fmt"

/*
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/

type ListNode struct {
	Val int
	Next *ListNode
}

// 初始化
func INIT(list []int) *ListNode {
	if list == nil || len(list) == 0 {
		return nil
	}
	head := new(ListNode)
	cur := head
	for i := 0; i < len(list); i++ {
		cur.Val = list[i]
		if i < (len(list) - 1) {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return head
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Printf("%v -> ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

// 非递归解法
func reverseGen(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 递归解法
func reverseD(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseD(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	testList := []int{1,2,3,4,5}
	PrintList(INIT(testList))
	PrintList(reverseGen(INIT(testList)))
	PrintList(reverseD(INIT(testList)))
}

