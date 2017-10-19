package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	step := head
	for nil != step.Next {
		step = step.Next
	}
	reverse(head, head.Next)
	head.Next = nil
	return step
}

func reverse(cur *ListNode, next *ListNode) {
	if nil == next {
		return
	} else {
		reverse(next, next.Next)
		next.Next = cur
	}
}
