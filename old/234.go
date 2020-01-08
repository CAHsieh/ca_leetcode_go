package main

func isPalindrome(head *ListNode) bool {
	if nil == head || nil == head.Next {
		return true
	}
	slow := head
	fast := head
	for nil != fast.Next && nil != fast.Next.Next {
		slow = slow.Next
		fast = fast.Next.Next
	}
	for nil != fast.Next {
		fast = fast.Next
	}
	reverse(slow.Next, slow.Next.Next)
	slow.Next.Next = nil

	for nil != fast {
		if fast.Val != head.Val {
			return false
		}
		fast = fast.Next
		head = head.Next
	}
	return true
}

// func reverse(cur *ListNode, next *ListNode) {
// 	if nil == next {
// 		return
// 	}
// 	reverse(next, next.Next)
// 	next.Next = cur
// }
