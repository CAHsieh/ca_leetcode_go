package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	slow := head
	fast := head
	var preSlow *ListNode

	count := 0
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		preSlow = slow
		slow = slow.Next
		count++
	}

	if head.Next == nil && count == 0 {
		if n == 0 {
			return head
		}
		return nil
	}

	var length int
	if fast.Next != nil {
		length = (count + 1) * 2
	} else {
		length = count*2 + 1
	}

	length = length - n

	if length < count {
		slow = head
		preSlow = nil
		for i := 0; i < length; i++ {
			preSlow = slow
			slow = slow.Next
		}
	} else if length > count {
		for ; count < length; count++ {
			preSlow = slow
			slow = slow.Next
		}
	}

	if preSlow != nil {
		preSlow.Next = slow.Next
	} else {
		return slow.Next
	}

	return head
}
